package json

import (
	"io"
	"sync"
	"sync/atomic"

	"go.melnyk.org/mlog"
)

type logbook struct {
	output       *syncwriter
	defaultlevel mlog.Level
	mu           sync.Mutex
	loggers      map[string]*logger
}

// Interface implementation check
var (
	_ mlog.Logbook = &logbook{}
)

func (lb *logbook) SetLevel(name string, level mlog.Level) error {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Set default level...
	if name == mlog.Default {
		lb.defaultlevel = level
		// so update all non-custom loggers too
		for _, v := range lb.loggers {
			if !v.customlevel {
				atomic.StoreUint32((*uint32)(&v.level), uint32(level))
			}
		}
	}

	// Set level for dedicated logger makes it custom
	if l, ok := lb.loggers[name]; ok {
		atomic.StoreUint32((*uint32)(&l.level), uint32(level))
		l.customlevel = true
		return nil
	}

	// Well...logger name is unknown, so add new one to logbook
	l := &logger{
		name:        name,
		out:         lb.output,
		level:       level,
		customlevel: true,
	}

	lb.loggers[name] = l

	return nil
}

func (lb *logbook) Levels() mlog.Levels {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	lvs := make(mlog.Levels)

	lvs[mlog.Default] = lb.defaultlevel
	for k, v := range lb.loggers {
		lvs[k] = mlog.Level(atomic.LoadUint32((*uint32)(&v.level)))
	}

	return lvs
}

func (lb *logbook) Joiner() mlog.Joiner {
	// logbook provides this interface in this implemenation
	return lb
}

func (lb *logbook) Join(name string) mlog.Logger {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	// Return exist logger
	if l, ok := lb.loggers[name]; ok {
		return l
	}

	// ... or create new one
	l := &logger{
		name:  name,
		out:   lb.output,
		level: lb.defaultlevel,
	}

	lb.loggers[name] = l

	return l
}

// NewLogbook returns interface to console logbook implementation
func NewLogbook(out io.Writer) mlog.Logbook {
	lb := &logbook{
		output:  &syncwriter{w: out},
		loggers: make(map[string]*logger),
	}

	return lb
}
