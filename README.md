# vegagoja

`vegagoja` renders [Vega visualizations][vega-examples] and [Vega-lite
visualizations][vega-lite-examples] as SVGs using the [`goja`][goja] JavaScript
runtime. Developed for use by [`usql`][usql] for rendering charts.

[Overview][] | [TODO][] | [About][]

[Overview]: #overview "Overview"
[TODO]: #todo "TODO"
[About]: #about "About"

[![Unit Tests][vegagoja-ci-status]][vegagoja-ci]
[![Go Reference][goref-vegagoja-status]][goref-vegagoja]
[![Discord Discussion][discord-status]][discord]

[vegagoja-ci]: https://github.com/xo/vegagoja/actions/workflows/test.yml
[vegagoja-ci-status]: https://github.com/xo/vegagoja/actions/workflows/test.yml/badge.svg
[goref-vegagoja]: https://pkg.go.dev/github.com/xo/vegagoja
[goref-vegagoja-status]: https://pkg.go.dev/badge/github.com/xo/vegagoja.svg
[discord]: https://discord.gg/yJKEzc7prt "Discord Discussion"
[discord-status]: https://img.shields.io/discord/829150509658013727.svg?label=Discord&logo=Discord&colorB=7289da&style=flat-square "Discord Discussion"

## Overview

Install in the usual Go fashion:

```sh
$ go get github.com/xo/vegagoja@latest
```

Then use like the following:

```go
package vegagoja_test

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/xo/vegagoja"
)

func Example() {
	vega := vegagoja.New(
		vegagoja.WithDemoData(),
	)
	data, err := vega.Render(context.Background(), candlestickSpec)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("candestick.svg", []byte(data), 0o644); err != nil {
		log.Fatal(err)
	}
	// Output:
}

const candlestickSpec = `{
  "$schema": "https://vega.github.io/schema/vega-lite/v5.json",
  "width": 400,
  "description": "A candlestick chart inspired by an example in Protovis (http://mbostock.github.io/protovis/ex/candlestick.html)",
  "data": {"url": "data/ohlc.json"},
  "encoding": {
    "x": {
      "field": "date",
      "type": "temporal",
      "title": "Date in 2009",
      "axis": {
        "format": "%m/%d",
        "labelAngle": -45,
        "title": "Date in 2009"
      }
    },
    "y": {
      "type": "quantitative",
      "scale": {"zero": false},
      "axis": {"title": "Price"}
    },
    "color": {
      "condition": {
        "test": "datum.open < datum.close",
        "value": "#06982d"
      },
      "value": "#ae1325"
    }
  },
  "layer": [
    {
      "mark": "rule",
      "encoding": {
        "y": {"field": "low"},
        "y2": {"field": "high"}
      }
    },
    {
      "mark": "bar",
      "encoding": {
        "y": {"field": "open"},
        "y2": {"field": "close"}
      }
    }
  ]
}`
```

## TODO

- Rewrite as native Go

## About

`vegagoja` was written primarily to support these projects:

- [usql][usql] - a universal command-line interface for SQL databases

Users of this package may find the [`github.com/xo/resvg`][resvg] package
helpful in rendering the

[usql]: https://github.com/xo/usql
[resvg]: https://github.com/xo/resvg
[goja]: https://github.com/dop251/goja
[vega]: https://vega.github.io
[vega-examples]: https://vega.github.io/vega/examples/
[vega-lite-examples]: https://vega.github.io/vega-lite/examples/
