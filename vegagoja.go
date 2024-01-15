// Package vegagoja renders Vega visualizations as SVGs.
package vegagoja

import (
	"context"
)

// DefaultVM is the default goja vm.
var DefaultVM = New()

// Version returns the embedded vega version.
func Version() (string, error) {
	return DefaultVM.Version()
}

// Render renders the spec with the specified data.
func Render(ctx context.Context, spec, data string) (string, error) {
	return DefaultVM.Render(ctx, spec, data)
}
