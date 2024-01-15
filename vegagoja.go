// Package vegagoja renders Vega visualizations as SVGs.
package vegagoja

import (
	"context"
	"embed"
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

// Vega handles rendering Vega visualizations as SVGs.
//
// Wraps a goja runtime vm, and uses embedded javascript to render the Vega
// visualizations.
type Vega struct {
	r           *goja.Runtime
	vegaVersion func() string
	render      func(logger func([]string), loader func(string) (string, error), cb func(string), spec string) string
	logger      func(...interface{})
	loader      func(string) ([]byte, error)
	data        fs.FS
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
		for _, name := range []string{"vega.min.js", "vegagoja.js"} {
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
		if err := r.ExportTo(r.Get("vega_version"), &vm.vegaVersion); err != nil {
			vm.err = fmt.Errorf("unable to export version func: %w", err)
			return
		}
		if err := r.ExportTo(r.Get("render"), &vm.render); err != nil {
			vm.err = fmt.Errorf("unable to export render func: %w", err)
			return
		}
		vm.r = r
	})
	return vm.err
}

// Version returns the embedded vega version.
func (vm *Vega) Version() (string, error) {
	if err := vm.init(); err != nil {
		return "", err
	}
	return strings.TrimPrefix(vm.vegaVersion(), "v"), nil
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
	ch := make(chan struct{})
	vm.render(vm.log, vm.load, func(s string) {
		defer close(ch)
		res = s
	}, spec)
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

// load is the script callback for loading a remote url.
func (vm *Vega) load(name string) (string, error) {
	switch {
	case vm.loader != nil:
		buf, err := vm.loader(name)
		if err != nil {
			return "", fmt.Errorf("loader could not open %s: %w", name, err)
		}
		return string(buf), nil
	case vm.data != nil:
		f, err := vm.data.Open(name)
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

// WithLoader is a vega option to set the loader.
func WithLoader(loader func(string) ([]byte, error)) Option {
	return func(vm *Vega) {
		vm.loader = loader
	}
}

// WithData is a vega option to set a [fs.FS] from which to read data.
func WithData(data fs.FS) Option {
	return func(vm *Vega) {
		vm.data = data
	}
}

// WithVegaDemoData is a vega option to add the embedded vega demo data. Useful
// for rendering Vega's example visualizations. Additional sources can be
// provided that will take priority when loading data.
func WithVegaDemoData(sources ...fs.FS) Option {
	return func(vm *Vega) {
		vm.data = &fallbackFS{
			sources: append(sources, vegaData),
		}
	}
}

// fallbackFS is a fallback [fs.FS] implementation.
type fallbackFS struct {
	sources []fs.FS
}

// Open satisfies the [fs.FS] interface.
func (f *fallbackFS) Open(name string) (fs.File, error) {
	for _, source := range f.sources {
		if file, err := source.Open(name); err == nil {
			return file, nil
		}
	}
	return nil, os.ErrNotExist
}

// vegaVersionTxt is the embedded vega-version.txt.
//
//go:embed vega-version.txt
var vegaVersionTxt string

// vegaScripts are the embedded vega javascripts.
//
//go:embed *.js
var vegaScripts embed.FS

// vegaData is the embedded vega data.
//
//go:embed data
var vegaData embed.FS
