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
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"strings"
	"sync"

	"github.com/dop251/goja"
	gojaparser "github.com/dop251/goja/parser"
	"github.com/dop251/goja_nodejs/console"
	"github.com/dop251/goja_nodejs/require"
)

// Vega handles rendering [Vega] and [Vega-Lite] visualizations as SVGs,
// wrapping a [goja] runtime vm.
//
// [Vega]: https://vega.github.io/vega/examples/
// [Vega-Lite]: https://vega.github.io/vega-lite/examples/
// [goja]: https://github.com/dop251/goja
type Vega struct {
	r       *goja.Runtime
	vegaVer func() string
	liteVer func() string
	render  renderFunc
	compile compileFunc
	logger  func(...interface{})
	source  fs.FS
	once    sync.Once
	err     error
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
		registry := new(require.Registry)
		r := goja.New()
		registry.Enable(r)
		console.Enable(r)
		for _, name := range []string{"vega.min.js", "vega-lite.min.js", "vegagoja.js"} {
			buf, err := vegaScripts.ReadFile(name)
			if err != nil {
				vm.err = fmt.Errorf("unable to load %s: %w", name, err)
				return
			}
			prg, err := goja.Parse(name, string(buf), gojaparser.WithDisableSourceMaps)
			if err != nil {
				vm.err = fmt.Errorf("unable to parse %s: %w", name, err)
				return
			}
			p, err := goja.CompileAST(prg, true)
			if err != nil {
				vm.err = fmt.Errorf("unable to compile %s: %w", name, err)
				return
			}
			if _, err := r.RunProgram(p); err != nil {
				vm.err = fmt.Errorf("unable to run %s: %w", name, err)
				return
			}
		}
		if err := r.ExportTo(r.Get("vega_version"), &vm.vegaVer); err != nil {
			vm.err = fmt.Errorf("unable to bind vega_version func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("lite_version"), &vm.liteVer); err != nil {
			vm.err = fmt.Errorf("unable to bind lite_version func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("render"), &vm.render); err != nil {
			vm.err = fmt.Errorf("unable to bind render func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("compile"), &vm.compile); err != nil {
			vm.err = fmt.Errorf("unable to bind compile func: %w", err)
			return
		}
		vm.r = r
	})
	return vm.err
}

// Version returns the embedded vega version.
func (vm *Vega) Version() (string, string, error) {
	if err := vm.init(); err != nil {
		return "", "", err
	}
	vegaVer := strings.TrimPrefix(vm.vegaVer(), "v")
	liteVer := strings.TrimPrefix(vm.liteVer(), "v")
	return vegaVer, liteVer, nil
}

// CompileSpec compiles a vega-lite specification to a vega specification,
// returning the entire raw compiled json containing the "spec" and
// "normalized" fields.
func (vm *Vega) CompileSpec(spec string) (string, error) {
	if err := vm.init(); err != nil {
		return "", err
	}
	return vm.compile(vm.log, spec)
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

// Render renders a vega visualization spec as a SVG.
func (vm *Vega) Render(ctx context.Context, spec string) (res string, err error) {
	if spec == "" {
		return
	}
	// unmarshal
	var m map[string]interface{}
	if err = json.Unmarshal([]byte(spec), &m); err != nil {
		err = ErrInvalidJSON
		return
	}
	// check schema
	s, ok := m["$schema"]
	if !ok {
		err = ErrMissingSchema
		return
	}
	schema, ok := s.(string)
	switch {
	case !ok:
		err = ErrSchemaInvalid
		return
	case !strings.HasPrefix(schema, "https://vega.github.io/schema/vega"):
		err = ErrNotVegaOrVegaLiteSchema
		return
	}
	// convert vega-lite -> vega
	if strings.HasPrefix(schema, "https://vega.github.io/schema/vega-lite/") {
		if spec, err = vm.Compile(spec); err != nil {
			err = fmt.Errorf("unable to compile vega-lite spec: %w", err)
			return
		}
	}
	// init
	if err = vm.init(); err != nil {
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
	ch, errch := make(chan struct{}, 1), make(chan error, 1)
	err = vm.render(vm.log, spec, vm.data(), func(s string) {
		defer close(ch)
		defer close(errch)
		res = s
	}, func(e string) {
		defer close(ch)
		defer close(errch)
		errch <- errors.New(e)
	})
	if err != nil {
		return
	}
	select {
	case <-ctx.Done():
		err = ctx.Err()
	case err = <-errch:
	case <-ch:
	}
	return
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

// WithDemoData is a vega option to add the vega demo data. Useful for
// rendering Vega's example visualizations. Additional source file systems
// ([fs.FS]) can be provided that will take priority when loading data.
func WithDemoData(sources ...fs.FS) Option {
	return func(vm *Vega) {
		vm.source = NewSource(append(sources, vegaData)...)
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
)

// Error satisfies the [error] interface.
func (err Error) Error() string {
	return string(err)
}

// loggerFunc is the signature for the log func.
type loggerFunc func([]string)

// renderFunc is the signature for the render func.
type renderFunc func(logf loggerFunc, spec string, data interface{}, cb, errcb func(string)) error

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

// vegaScripts are the embedded vega javascripts.
//
//go:embed *.js
var vegaScripts embed.FS

// vegaData is the embedded vega data.
//
//go:embed data
var vegaData embed.FS
