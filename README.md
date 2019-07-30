# mlog
[![GoDoc](https://godoc.org/go.melnyk.org/mlog?status.svg)](https://godoc.org/go.melnyk.org/mlog)

mlog is set of defined interfaces to add customable and high performance logging. It is not a logging framework with fancy output, but it allows to build this framework in effective way (see console logger as example).

## Why?

1. Unification of logging through different components
2. Effective implementation
3. Customization for main app needs

## How?

TBD

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
	lb := console.NewLogbook(os.Stdout)
	lb.SetLevel(mlog.Default, mlog.Verbose)
	logger:=lb.Joiner().Join("test")
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
