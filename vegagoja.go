// Package vegagoja renders Vega visualizations as SVGs.
package vegagoja

import (
	"context"
)

// DefaultVM is the default goja vm.
var DefaultVM = New()

// Version returns the embedded vega and vega lite versions.
func Version() (string, string, error) {
	vegaVer, err := DefaultVM.VegaVersion()
	if err != nil {
		return "", "", err
	}
	vegaLiteVer, err := DefaultVM.VegaLiteVersion()
	if err != nil {
		return "", "", err
	}
	return vegaVer, vegaLiteVer, nil
}

// Render renders the spec with the specified data.
func Render(ctx context.Context, spec, data string) (string, error) {
	return DefaultVM.Render(ctx, spec, data)
}
