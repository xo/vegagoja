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
	"log"
	"os"

	"github.com/xo/vegagoja"
)

func Example() {
	const spec = `{
  "$schema": "https://vega.github.io/schema/vega/v5.json",
  "description": "A candlestick chart inspired by an example in Protovis (http://mbostock.github.io/protovis/ex/candlestick.html)",
  "background": "white",
  "padding": 5,
  "width": 400,
  "height": 200,
  "style": "cell",
  "data": [
    {
      "name": "source_0",
      "url": "data/ohlc.json",
      "format": {"type": "json", "parse": {"date": "date"}}
    },
    {
      "name": "data_0",
      "source": "source_0",
      "transform": [
        {
          "type": "filter",
          "expr": "(isDate(datum[\"date\"]) || (isValid(datum[\"date\"]) && isFinite(+datum[\"date\"]))) && isValid(datum[\"low\"]) && isFinite(+datum[\"low\"])"
        }
      ]
    },
    {
      "name": "data_1",
      "source": "source_0",
      "transform": [
        {
          "type": "filter",
          "expr": "(isDate(datum[\"date\"]) || (isValid(datum[\"date\"]) && isFinite(+datum[\"date\"]))) && isValid(datum[\"open\"]) && isFinite(+datum[\"open\"])"
        }
      ]
    }
  ],
  "marks": [
    {
      "name": "layer_0_marks",
      "type": "rule",
      "style": ["rule"],
      "from": {"data": "data_0"},
      "encode": {
        "update": {
          "stroke": [
            {"test": "datum.open < datum.close", "value": "#06982d"},
            {"value": "#ae1325"}
          ],
          "description": {
            "signal": "\"Date in 2009: \" + (timeFormat(datum[\"date\"], '%m/%d')) + \"; low: \" + (format(datum[\"low\"], \"\")) + \"; high: \" + (format(datum[\"high\"], \"\"))"
          },
          "x": {"scale": "x", "field": "date"},
          "y": {"scale": "y", "field": "low"},
          "y2": {"scale": "y", "field": "high"}
        }
      }
    },
    {
      "name": "layer_1_marks",
      "type": "rect",
      "style": ["bar"],
      "from": {"data": "data_1"},
      "encode": {
        "update": {
          "fill": [
            {"test": "datum.open < datum.close", "value": "#06982d"},
            {"value": "#ae1325"}
          ],
          "ariaRoleDescription": {"value": "bar"},
          "description": {
            "signal": "\"Date in 2009: \" + (timeFormat(datum[\"date\"], '%m/%d')) + \"; open: \" + (format(datum[\"open\"], \"\")) + \"; close: \" + (format(datum[\"close\"], \"\"))"
          },
          "xc": {"scale": "x", "field": "date"},
          "width": {"value": 5},
          "y": {"scale": "y", "field": "open"},
          "y2": {"scale": "y", "field": "close"}
        }
      }
    }
  ],
  "scales": [
    {
      "name": "x",
      "type": "time",
      "domain": {
        "fields": [
          {"data": "data_0", "field": "date"},
          {"data": "data_1", "field": "date"}
        ]
      },
      "range": [0, {"signal": "width"}],
      "padding": 5
    },
    {
      "name": "y",
      "type": "linear",
      "domain": {
        "fields": [
          {"data": "data_0", "field": "low"},
          {"data": "data_0", "field": "high"},
          {"data": "data_1", "field": "open"},
          {"data": "data_1", "field": "close"}
        ]
      },
      "range": [{"signal": "height"}, 0],
      "zero": false,
      "nice": true
    }
  ],
  "axes": [
    {
      "scale": "x",
      "orient": "bottom",
      "gridScale": "y",
      "grid": true,
      "tickCount": {"signal": "ceil(width/40)"},
      "domain": false,
      "labels": false,
      "aria": false,
      "maxExtent": 0,
      "minExtent": 0,
      "ticks": false,
      "zindex": 0
    },
    {
      "scale": "y",
      "orient": "left",
      "gridScale": "x",
      "grid": true,
      "tickCount": {"signal": "ceil(height/40)"},
      "domain": false,
      "labels": false,
      "aria": false,
      "maxExtent": 0,
      "minExtent": 0,
      "ticks": false,
      "zindex": 0
    },
    {
      "scale": "x",
      "orient": "bottom",
      "grid": false,
      "title": "Date in 2009",
      "format": "%m/%d",
      "labelAngle": 315,
      "labelAlign": "right",
      "labelBaseline": "top",
      "labelFlush": true,
      "labelOverlap": true,
      "tickCount": {"signal": "ceil(width/40)"},
      "zindex": 0
    },
    {
      "scale": "y",
      "orient": "left",
      "grid": false,
      "title": "Price",
      "labelOverlap": true,
      "tickCount": {"signal": "ceil(height/40)"},
      "zindex": 0
    }
  ]
}`
	vega := vegagoja.New(
		vegagoja.WithVegaDemoData(),
	)
	data, err := vega.Render(context.Background(), spec)
	if err != nil {
		log.Fatal(err)
	}
	if err := os.WriteFile("candestick.svg", []byte(data), 0o644); err != nil {
		log.Fatal(err)
	}
	// Output:
}
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
