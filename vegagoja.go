// Package vegagoja renders [Vega] and [Vega-Lite] visualizations as SVGs using
// the [goja] JavaScript runtime. Developed for use by [usql] for rendering
// charts.
//
// [Vega]: https://vega.github.io/vega/examples/
// [Vega-Lite]: https://vega.github.io/vega-lite/examples/
// [goja]: https://github.com/dop251/goja
// [usql]: https://github.com/xo/usql
package vegagoja

import (
	"bytes"
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	"sync"

	"github.com/dop251/goja"
	gojaparser "github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/eventloop"
	"github.com/dop251/goja_nodejs/require"
)

// Vega handles rendering [Vega] and [Vega-Lite] visualizations as SVGs,
// wrapping a [goja] runtime vm.
//
// [Vega]: https://vega.github.io/vega/examples/
// [Vega-Lite]: https://vega.github.io/vega-lite/examples/
// [goja]: https://github.com/dop251/goja
type Vega struct {
	logger     func(...interface{})
	vegaJs     *goja.Program
	vegaLiteJs *goja.Program
	vegagojaJs *goja.Program
	source     fs.FS
	once       sync.Once
	err        error
}

// New creates a new vega instance.
func New(opts ...Option) *Vega {
	vm := new(Vega)
	for _, o := range opts {
		o(vm)
	}
	return vm
}

// init initializes the vega instance.
func (vm *Vega) init() error {
	vm.once.Do(func() {
		if vm.vegaJs, vm.err = compileEmbeddedScript("vega.min.js"); vm.err != nil {
			return
		}
		if vm.vegaLiteJs, vm.err = compileEmbeddedScript("vega-lite.min.js"); vm.err != nil {
			return
		}
		if vm.vegagojaJs, vm.err = compileEmbeddedScript("vegagoja.js"); vm.err != nil {
			return
		}
	})
	return vm.err
}

// run runs the embedded javascripts on the goja runtime, and exports a symbol
// name to v.
func (vm *Vega) run(r *goja.Runtime, name string, v interface{}) error {
	if _, err := r.RunProgram(vm.vegaJs); err != nil {
		return err
	}
	if _, err := r.RunProgram(vm.vegaLiteJs); err != nil {
		return err
	}
	if _, err := r.RunProgram(vm.vegagojaJs); err != nil {
		return err
	}
	if name != "" {
		return r.ExportTo(r.Get(name), v)
	}
	return nil
}

// loop instantiates a new loop, and waits for it to terminate, or until the
// context is closed.
func (vm *Vega) loop(ctx context.Context, f func(*goja.Runtime) error) error {
	// init
	if err := vm.init(); err != nil {
		return err
	}
	mod := console.RequireWithPrinter(vm)
	registry := new(require.Registry)
	registry.RegisterNativeModule(console.ModuleName, mod)
	loop := eventloop.NewEventLoop(
		eventloop.WithRegistry(registry),
	)
	errch := make(chan error, 1)
	go func() {
		loop.Run(func(r *goja.Runtime) {
			defer close(errch)
			errch <- f(r)
		})
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case err := <-errch:
		return err
	}
}

// Version returns the embedded vega version.
func (vm *Vega) Version() (string, string, error) {
	if err := vm.init(); err != nil {
		return "", "", err
	}
	r := goja.New()
	var f func() ([]string, error)
	if err := vm.run(r, "version", &f); err != nil {
		return "", "", err
	}
	ver, err := f()
	switch {
	case err != nil:
		return "", "", err
	case len(ver) != 2:
		return "", "", ErrInvalidResult
	}
	return ver[0], ver[1], nil
}

// CompileSpec compiles a vega-lite specification to a vega specification,
// returning the entire raw compiled json containing the "spec" and
// "normalized" fields.
func (vm *Vega) CompileSpec(spec string) (string, error) {
	if err := vm.init(); err != nil {
		return "", err
	}
	var compile compileFunc
	if err := vm.run(goja.New(), "compile", &compile); err != nil {
		return "", err
	}
	return compile(vm.log, spec)
}

// Compile compiles a vega-lite specification to a vega specification.
//
// Wraps [CompileSpec], returning only the compiled "spec".
func (vm *Vega) Compile(spec string) (string, error) {
	spec, err := vm.CompileSpec(spec)
	if err != nil {
		return "", err
	}
	var res map[string]interface{}
	if err := json.Unmarshal([]byte(spec), &res); err != nil {
		return "", ErrInvalidCompiledSpec
	}
	s, ok := res["spec"]
	if !ok {
		return "", ErrInvalidCompiledSpec
	}
	buf := new(bytes.Buffer)
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(false)
	if err := enc.Encode(s); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// CheckSpec takes a spec and checks that it is a vega or vega-lite spec. If it
// is a vega-lite spec, it will returned the compiled vega spec.
func (vm *Vega) CheckSpec(spec string) (string, error) {
	// unmarshal
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(spec), &m); err != nil {
		return "", ErrInvalidJSON
	}
	// check schema
	s, ok := m["$schema"]
	if !ok {
		return "", ErrMissingSchema
	}
	schema, ok := s.(string)
	switch {
	case !ok:
		return "", ErrSchemaInvalid
	case !strings.HasPrefix(schema, "https://vega.github.io/schema/vega"):
		return "", ErrNotVegaOrVegaLiteSchema
	case strings.HasPrefix(schema, "https://vega.github.io/schema/vega/"):
		return spec, nil
	}
	// convert vega-lite -> vega
	spec, err := vm.Compile(spec)
	if err != nil {
		return "", fmt.Errorf("unable to compile vega-lite spec: %w", err)
	}
	return spec, nil
}

// Render renders a vega visualization spec as a SVG.
func (vm *Vega) Render(ctx context.Context, spec string) (res string, err error) {
	if spec, err = vm.CheckSpec(spec); err != nil {
		return
	}
	defer func() {
		if e := recover(); e != nil {
			if ex, ok := e.(*goja.Exception); ok {
				res, err = "", ex.Unwrap()
			} else {
				res, err = "", fmt.Errorf("recovered from: %v", e)
			}
		}
	}()
	err = vm.loop(ctx, func(r *goja.Runtime) error {
		var f renderFunc
		if err := vm.run(r, "render", &f); err != nil {
			return err
		}
		promise := f(vm.log, spec, vm.data())
		switch state := promise.State(); state {
		case goja.PromiseStateFulfilled:
			result := promise.Result()
			if result == nil {
				return ErrInvalidResult
			}
			res = result.String()
		case goja.PromiseStateRejected:
			result := promise.Result()
			if obj, ok := result.(*goja.Object); ok {
				if stack := obj.Get("stack"); stack != nil {
					return fmt.Errorf("caught error during rendering: %s", stack.String())
				}
			}
			return fmt.Errorf("unknown rejected promise result: %v", result)
		default:
			return fmt.Errorf("unknown promise state: %v", state)
		}
		return nil
	})
	return
}

// Log satisfies the [goja.Printer] interface.
func (vm *Vega) Log(s string) {
	vm.log([]string{"LOG", s})
}

// Warn satisfies the [goja.Printer] interface.
func (vm *Vega) Warn(s string) {
	vm.log([]string{"WARN", s})
}

// Error satisfies the [goja.Printer] interface.
func (vm *Vega) Error(s string) {
	vm.log([]string{"ERROR", s})
}

// log is the script callback for logging a message.
func (vm *Vega) log(s []string) {
	if vm.logger != nil {
		v := make([]interface{}, len(s))
		for i, ss := range s {
			v[i] = ss
		}
		vm.logger(v...)
	}
}

// data returns the data.
func (vm *Vega) data() interface{} {
	return vm.load
}

// load loads data from sources.
func (vm *Vega) load(name string) (string, error) {
	if vm.source != nil {
		f, err := vm.source.Open(name)
		if err != nil {
			return "", fmt.Errorf("could not open from data %s: %w", name, err)
		}
		defer f.Close()
		buf, err := io.ReadAll(f)
		if err != nil {
			return "", fmt.Errorf("unable to read all data for %s: %w", name, err)
		}
		return string(buf), nil
	}
	return "", fmt.Errorf("no loader for %s: %w", name, os.ErrNotExist)
}

// Option is a vega option.
type Option func(*Vega)

// WithLogger is a vega option to set the logger.
func WithLogger(logger func(...interface{})) Option {
	return func(vm *Vega) {
		vm.logger = logger
	}
}

// WithResultSet is a vega option to set the data from a result set.
func WithResultSet(resultSet ResultSet) Option {
	return func(vm *Vega) {
	}
}

// WithRecords is a vega option to set the data from a set of headers and
// records.
func WithRecords(headers []string, records [][]string) Option {
	return func(vm *Vega) {
	}
}

// WithCSV is a vega option to read csv data from the supplied reader.
func WithCSV(r io.Reader) Option {
	return func(vm *Vega) {
	}
}

// WithCSVString is a vega option to read csv data from the string.
func WithCSVString(s string) Option {
	return func(vm *Vega) {
		WithCSV(strings.NewReader(s))(vm)
	}
}

// WithCSVBytes is a vega option to read csv data from the bytes.
func WithCSVBytes(buf []byte) Option {
	return func(vm *Vega) {
		WithCSV(bytes.NewReader(buf))(vm)
	}
}

// WithSources is a vega option to set the source file systems ([fs.FS]) from
// which to load data.
func WithSources(sources ...fs.FS) Option {
	return func(vm *Vega) {
		vm.source = NewSource(sources...)
	}
}

// WithDataDir is a vega option to add a data source that loads data from a
// directory.
func WithDataDir(dir string) Option {
	return func(vm *Vega) {
		vm.source = os.DirFS(dir)
	}
}

// WithPrefixedSourceDir is a vega option to add a data source that loads data
// from a specified prefixed directory name.
func WithPrefixedSourceDir(prefix, dir string) Option {
	return func(vm *Vega) {
		vm.source = NewPrefixedSourceDir(prefix, dir)
	}
}

// Error is a error.
type Error string

// Errors.
const (
	// ErrInvalidCompiledSpec is the invalid compiled spec error.
	ErrInvalidCompiledSpec Error = "invalid compiled spec"
	// ErrInvalidJSON is the invalid json error.
	ErrInvalidJSON Error = "invalid json"
	// ErrMissingSchema is the missing $schema error.
	ErrMissingSchema Error = "missing $schema"
	// ErrSchemaInvalid is the $schema invalid error.
	ErrSchemaInvalid Error = "$schema invalid"
	// ErrNotVegaOrVegaLiteSchema is the not vega or vega-lite schema error.
	ErrNotVegaOrVegaLiteSchema Error = "not vega or vega-lite schema"
	// ErrInvalidResult is the invalid result error.
	ErrInvalidResult Error = "invalid result"
)

// Error satisfies the [error] interface.
func (err Error) Error() string {
	return string(err)
}

// compile compiles a goja program.
func compile(name, src string) (*goja.Program, error) {
	prg, err := goja.Parse(name, src, gojaparser.WithDisableSourceMaps)
	if err != nil {
		return nil, fmt.Errorf("unable to parse %s: %w", name, err)
	}
	p, err := goja.CompileAST(prg, true)
	if err != nil {
		return nil, fmt.Errorf("unable to compile %s: %w", name, err)
	}
	return p, nil
}

// compileEmbeddedScript compiles the embedded script as a goja program.
func compileEmbeddedScript(name string) (*goja.Program, error) {
	buf, err := jsScripts.ReadFile(name)
	if err != nil {
		return nil, fmt.Errorf("unable to load %s: %w", name, err)
	}
	return compile(name, string(buf))
}

// loggerFunc is the signature for the log func.
type loggerFunc func([]string)

// renderFunc is the signature for the render func.
type renderFunc func(logf loggerFunc, spec string, data interface{}) *goja.Promise

// compileFunc is the signature for the compile func.
type compileFunc func(logf loggerFunc, spec string) (string, error)

// vegaVersionTxt is the embedded vega-version.txt.
//
//go:embed vega-version.txt
var vegaVersionTxt string

// liteVersionTxt is the embedded vega-lite-version.txt.
//
//go:embed vega-lite-version.txt
var liteVersionTxt string

// jsScripts are the embedded vega javascripts.
//
//go:embed *.js
var jsScripts embed.FS
