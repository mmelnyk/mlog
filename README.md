# mlog
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://github.com/mmelnyk/mlog/blob/master/LICENSE) [![Actions Status](https://github.com/mmelnyk/mlog/workflows/test/badge.svg)](https://github.com/mmelnyk/mlog/actions)[![GoDoc](https://godoc.org/go.melnyk.org/mlog?status.svg)](https://godoc.org/go.melnyk.org/mlog) [![Go Report Card](https://goreportcard.com/badge/go.melnyk.org/mlog)](https://goreportcard.com/report/go.melnyk.org/mlog)

mlog is set of defined interfaces to add customable and high performance logging. It is not a logging framework with fancy output, but it allows to build this framework in effective way (see console logger as example).

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
- Very simple
- Flexibility and High level of customization

## How?
## Flexibility
Main mlog module just defines set of interfaces for logbooks and loggers. Also mlog includes basic (but performant) nolog and console inmplementations. In case if special logger needed (e.g. structured), it can be easly implemented (mlog has very simple interface).

## Performance
Basic interface and logging via closires allow mlog to have very minimal instrumentation overhead and zero allocation logging. See benchmarks.

### Leveled Logging
mlog defines logging at the following levels:
 - Fatal
 - Error
 - Warning
 - Info
 - Verbose

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
BenchmarkDisabled/sirupsen/logrus-6         	50000000	        21.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkDisabled/uber-go/zap-6             	100000000	        12.1 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabled/rs/zerolog-6              	1000000000	         2.42 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabled/bloom42/rz-go-6           	1000000000	         1.77 ns/op	       0 B/op	       0 allocs/op
BenchmarkDisabled/mmelnyk/mlog-6            	2000000000	         0.82 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticText/sirupsen/logrus-6       	  200000	      7048 ns/op	     437 B/op	      16 allocs/op
BenchmarkStaticText/uber-go/zap-6           	 3000000	       697 ns/op	      72 B/op	       3 allocs/op
BenchmarkStaticText/rs/zerolog-6            	 5000000	       331 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticText/bloom42/rz-go-6         	 5000000	       324 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticText/mmelnyk/mlog-6          	10000000	       299 ns/op	       0 B/op	       0 allocs/op
BenchmarkStaticText/mmelnyk/mlog/c-6        	10000000	       214 ns/op	       0 B/op	       0 allocs/op
```
