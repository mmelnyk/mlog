# :ledger: mlog
[![License][license-img]][license] [![Actions Status][action-img]][action] [![GoDoc][godoc-img]][godoc] [![Go Report Card][goreport-img]][goreport] [![Coverage Status][codecov-img]][codecov]

mlog is set of defined interfaces to add customable and high performance logging. It is not a logger with fancy output, but it allows to build any logger in effective way using mlog as base framework (see console logger as example).

## Why?
We can find many fantastic and popular loggers (just a few examples - logrus and zap) for our golang projects and they work very well in most condition. However, there is a few cases when they do not work as expected:
 - logger for "as performant as possible" code - yes, logs are good for troubleshooting, but logger can introduce huge performance impact on our highly optimized code.
 - "output polution" - have you used 3rd party modules with enabled logging in your project?

So, I created mlog to solve following tasks:
1. Unification of logging through different components
2. Framework for effective logging implementation
3. Customization for main app needs

## Features
- Blazing fast
- Low to zero allocation
- Level logging
- Very
- No dependencies
- Flexibility and High level of customization

## How?
## Flexibility
Main mlog module just defines set of interfaces for logbooks and loggers. Also mlog includes basic (but performant) nolog and console inmplementations. In case if special logger needed (e.g. structured), it can be easly implemented (mlog has very simple interface).

## Performance
Basic interface and logging via closires allow mlog to have very minimal instrumentation overhead and zero allocation logging. See benchmarks.

### Leveled Logging
mlog defines logging at the following levels:
 - Fatal
 - Panic
 - Error
 - Warning
 - Info
 - Verbose

:exclamation: **Warning** - logger.Fatal and logger.Panic do **NOT DO** calls panic() or os.Exit() functions. You **MUST** consider **the logger does not control your application flow**.

We can control default and dedicasted logging level via logbook's SetLevel call.
Example:
```
...
	lb := console.NewLogbook(os.Stdout)
	lb.SetLevel(mlog.Default, mlog.Error) // Set Error as default level for whole logbook
	lb.SetLevel("myapp", mlog.Verbose)    // Set Verbose for myapp logger
...
```

## Concept
mlog defines:
- Logbook as main holder for all loggers
- Joiner as a way to get/join logger into logbook
- Logger as a way to output log messages

## Avalable implementations

### Official
 * console - pretty console-formated logger
 * nolog - performance optimized output supressed logger
 * testlog - silent logger for unittests and code coverage tests

### Additional
Please check [https://github.com/mlog-adapters](https://github.com/mlog-adapters).

## Examples

### Simple console logger
```
import (
	...
	"os"
	...
	"go.melnyk.org/mlog"
	"go.melnyk.org/mlog/console"
	...
)

...
	lb := console.NewLogbook(os.Stdout) // Create new logbook from console implementation
	lb.SetLevel(mlog.Default, mlog.Verbose) // Set default logging level to Verbose
	logger:=lb.Joiner().Join("test") // Get (join) logger with name "test"
...
	// Direct call
	logger.Info("Info message")

	// Via closure callback
	logger.Event(mlog.Info, func(ev mlog.Event) {
		ev.String("msg", "Info message")
	})
...

```

## Benchmarks

```
pkg: go.melnyk.org/mlog/benchmarks
BenchmarkDisabled/sirupsen/logrus-36       138527689    9.48 ns/op      16 B/op	       1 allocs/op
BenchmarkDisabled/uber-go/zap-36          1000000000   0.747 ns/op       0 B/op	       0 allocs/op
BenchmarkDisabled/rs/zerolog-36           1000000000   0.346 ns/op       0 B/op	       0 allocs/op
BenchmarkDisabled/mmelnyk/mlog-36         1000000000   0.135 ns/op       0 B/op	       0 allocs/op
BenchmarkStaticText/sirupsen/logrus-36        311452    3658 ns/op     456 B/op	      16 allocs/op
BenchmarkStaticText/uber-go/zap-36           5399784     230 ns/op      73 B/op	       3 allocs/op
BenchmarkStaticText/rs/zerolog-36            4862523     233 ns/op       0 B/op	       0 allocs/op
BenchmarkStaticText/mmelnyk/mlog-36          5122460     225 ns/op       0 B/op	       0 allocs/op
BenchmarkStaticText/mmelnyk/mlog/c-36        5338666     234 ns/op       0 B/op        0 allocs/op
Benchmark10Fields/sirupsen/logrus-36           80421   14605 ns/op    3021 B/op       47 allocs/op
Benchmark10Fields/uber-go/zap-36             1359621     892 ns/op     732 B/op        4 allocs/op
Benchmark10Fields/rs/zerolog-36              4673356     258 ns/op       0 B/op        0 allocs/op
Benchmark10Fields/mmelnyk/mlog/c-36          4922166     253 ns/op       0 B/op        0 allocs/op
```

## Development Status: Stable
All APIs are finalized, and no breaking changes will be made in the 1.x series
of releases.


[license-img]: https://img.shields.io/badge/license-MIT-blue.svg
[license]: https://github.com/mmelnyk/mlog/blob/master/LICENSE
[action-img]: https://github.com/mmelnyk/mlog/workflows/Test/badge.svg
[action]: https://github.com/mmelnyk/mlog/actions
[godoc-img]: https://godoc.org/go.melnyk.org/mlog?status.svg
[godoc]: https://godoc.org/go.melnyk.org/mlog
[goreport-img]: https://goreportcard.com/badge/go.melnyk.org/mlog
[goreport]: https://goreportcard.com/report/go.melnyk.org/mlog
[codecov-img]: https://codecov.io/gh/mmelnyk/mlog/branch/master/graph/badge.svg
[codecov]: https://codecov.io/gh/mmelnyk/mlog
