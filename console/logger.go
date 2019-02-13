package console

import (
	"io"
	"runtime"
	"sync/atomic"

	"go.melnyk.org/mlog"
)

type logger struct {
	name        string
	customlevel bool
	level       mlog.Level
	out         io.Writer
}

func (l *logger) Verbose(msg string) {
	l.output(mlog.Verbose, msg, nil)
}
func (l *logger) Info(msg string) {
	l.output(mlog.Info, msg, nil)
}
func (l *logger) Warning(msg string) {
	l.output(mlog.Warning, msg, nil)
}
func (l *logger) Error(msg string) {
	l.output(mlog.Error, msg, nil)
}
func (l *logger) Fatal(msg string) {
	l.output(mlog.Fatal, msg, nil)
}
func (l *logger) Event(level mlog.Level, cb func(evt mlog.Event)) {
	l.output(level, "", cb)
}

func (l *logger) output(level mlog.Level, str string, cb func(evt mlog.Event)) {
	// Level filter
	if lv := mlog.Level(atomic.LoadUint32((*uint32)(&l.level))); lv < level || lv == mlog.None {
		return
	}

	evt := getEvent()
	evt.buffer.Reset()

	// Header
	evt.addTimestamp()
	//evt.String(fieldLevel, l.levelToString(level)) // for structured log
	evt.justJoinString(l.levelDisplay(level)) // just simple console log
	//evt.String(fieldName, l.name) // for structured log
	evt.justJoinString(l.name) // just simple console log

	// Message or custom event
	if cb != nil {
		cb(evt)
	} else {
		evt.justJoinString(str)
	}

	// Footer

	// Location of error event caller
	if level == mlog.Error {
		var pcsb [20]uintptr
		pcs := pcsb[:]
		depth := runtime.Callers(3, pcs)
		frames := runtime.CallersFrames(pcs[:depth])
		if f, again := frames.Next(); again {
			evt.singleframe(&f)
		}
	}

	evt.buffer.WriteByte('\n')

	// Callstack for fatal event
	if level == mlog.Fatal {
		var pcsb [20]uintptr
		pcs := pcsb[:]
		depth := runtime.Callers(3, pcs)
		frames := runtime.CallersFrames(pcs[:depth])

		i := 1
		for f, again := frames.Next(); again; f, again = frames.Next() {
			evt.frame(i, &f)
			i++
		}
	}

	// Flush buffer to the logbook stream
	l.out.Write(evt.buffer.Bytes())

	// Return event to the pool
	putEvent(evt)
}

func (l *logger) levelDisplay(lv mlog.Level) (name string) {
	switch lv {
	case mlog.None:
		name = "NONE    "
	case mlog.Fatal:
		name = "\033[31mFATAL   \033[0m"
	case mlog.Error:
		name = "\033[31mERROR   \033[0m"
	case mlog.Warning:
		name = "\033[33mWARNING \033[0m"
	case mlog.Info:
		name = "\033[32mINFO    \033[0m"
	case mlog.Verbose:
		name = "VERBOSE "
	default:
		name = "UNKNOWN"
	}
	return
}
