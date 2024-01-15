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
