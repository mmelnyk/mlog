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

// Interface implementation check
var (
	_ mlog.Logger = &logger{}
)

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
	if lv := mlog.Level(atomic.LoadUint32((*uint32)(&l.level))); lv < level {
		return
	}

	evt := getEvent()

	// Header
	evt.addTimestamp()
	evt.justJoinLevel(level)   // just simple console log
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
