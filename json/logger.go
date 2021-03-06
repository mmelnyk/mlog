package json

import (
	"io"
	"runtime"

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
func (l *logger) Panic(msg string) {
	l.output(mlog.Panic, msg, nil)
}
func (l *logger) Fatal(msg string) {
	l.output(mlog.Fatal, msg, nil)
}
func (l *logger) Event(level mlog.Level, cb func(evt mlog.Event)) {
	l.output(level, "", cb)
}

func (l *logger) output(level mlog.Level, str string, cb func(evt mlog.Event)) {
	// Level filter
	// why not atomic.LoadUint32... it's not expected the level is changed dynamicly
	// too often, and even we have some missed sync between CPU's L1 caches, it's not
	// critical for logging. So.. don't lock CPU ;)
	if l.level < level {
		return
	}

	evt := getEvent()

	// Header
	evt.header(level, l.name) // just simple console log

	// Message or custom event
	if cb != nil {
		cb(evt)
	} else {
		evt.String(fieldMessage, str)
	}

	// Location of error event caller
	if level == mlog.Error {
		var pcsb [20]uintptr
		pcs := pcsb[:]
		depth := runtime.Callers(3, pcs)
		frames := runtime.CallersFrames(pcs[:depth])
		if f, again := frames.Next(); again {
			evt.buffer.WriteString(`"code":`)
			evt.frame(&f)
			evt.buffer.WriteString(", ")
		}
	}

	// Callstack for fatal event
	if level <= mlog.Panic {
		var pcsb [20]uintptr
		pcs := pcsb[:]
		depth := runtime.Callers(3, pcs)
		frames := runtime.CallersFrames(pcs[:depth])

		evt.buffer.WriteString(`"stack":[`)

		sep := false
		for f, again := frames.Next(); again; f, again = frames.Next() {
			if sep {
				evt.buffer.WriteString(", ")
			} else {
				sep = true
			}
			evt.frame(&f)
		}
		evt.buffer.WriteString(`], `)
	}

	// Add tail
	evt.tail()

	// Flush buffer to the logbook stream
	l.out.Write(evt.buffer.Bytes())

	// Return event to the pool
	putEvent(evt)
}
