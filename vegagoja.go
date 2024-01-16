// Package vegagoja renders Vega and Vega Lite visualizations as SVGs.
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
	"github.com/dop251/goja_nodejs/require"
)

// Vega handles rendering Vega and Vega Lite visualizations as SVGs. Wraps a
// goja runtime vm, and uses embedded javascript to render Vega and Vega Lite
// visualizations.
type Vega struct {
	r           *goja.Runtime
	vegaVer     func() string
	liteVer     func() string
	vegaRender  renderFunc
	liteRender  renderFunc
	liteCompile func(loggerFunc, string) (string, error)
	logger      func(...interface{})
	source      fs.FS
	once        sync.Once
	err         error
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
				vm.err = fmt.Errorf("unable to load %s: %w", name, err)
				return
			}
		}
		if err := r.ExportTo(r.Get("vega_version"), &vm.vegaVer); err != nil {
			vm.err = fmt.Errorf("unable to bind vega_version func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("vega_lite_version"), &vm.liteVer); err != nil {
			vm.err = fmt.Errorf("unable to bind vega_lite_version func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("vega_render"), &vm.vegaRender); err != nil {
			vm.err = fmt.Errorf("unable to bind vega_render func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("vega_lite_compile"), &vm.liteCompile); err != nil {
			vm.err = fmt.Errorf("unable to bind vega_lite_compile func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("vega_lite_render"), &vm.liteRender); err != nil {
			vm.err = fmt.Errorf("unable to bind vega_lite_render func: %w", err)
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

// Compile compiles a vega lite specification to a vega specification.
func (vm *Vega) Compile(spec string) (string, error) {
	if err := vm.init(); err != nil {
		return "", err
	}
	return vm.liteCompile(vm.log, spec)
}

// Render renders the spec with the specified data.
func (vm *Vega) Render(ctx context.Context, spec string) (res string, err error) {
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
	f := vm.vegaRender
	switch s, ok := decodeSchema(spec); {
	case !ok:
		err = fmt.Errorf("spec does not contain or has invalid $schema definition")
		return
	case strings.Contains(s, "vega-lite"):
		f = vm.liteRender
	}
	ch := make(chan struct{})
	f(vm.log, spec, vm.data(), func(s string) {
		defer close(ch)
		res = s
	})
	select {
	case <-ctx.Done():
		err = ctx.Err()
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

// data returns the
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

// decodeSchema decodes the schema from the spec.
func decodeSchema(spec string) (string, bool) {
	// check $schema definition
	var m map[string]interface{}
	if err := json.NewDecoder(strings.NewReader(spec)).Decode(&m); err != nil {
		return "", false
	}
	s, ok := m["$schema"]
	if !ok {
		return "", false
	}
	schema, ok := s.(string)
	if !ok {
		return "", false
	}
	return schema, strings.HasPrefix(schema, vegaSchemaPrefix) || strings.HasPrefix(schema, liteSchemaPrefix)
}

// loggerFunc is the signature for the log func.
type loggerFunc func([]string)

// renderFunc is the signature for the render func.
type renderFunc func(logger loggerFunc, spec string, data interface{}, cb func(string)) string

// vega schema prefixes.
const (
	vegaSchemaPrefix = "https://vega.github.io/schema/vega/"
	liteSchemaPrefix = "https://vega.github.io/schema/vega-lite/"
)

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
