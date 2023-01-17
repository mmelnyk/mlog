package nolog

import (
	"go.melnyk.org/mlog"
)

type logger struct {
	// nolog logger
}

func (*logger) Verbose(string) {}
func (*logger) Info(string)    {}
func (*logger) Warning(string) {}
func (*logger) Error(string)   {}
func (*logger) Panic(string)   {}
func (*logger) Fatal(string)   {}

func (*logger) Event(level mlog.Level, cb func(lg mlog.Event)) {}

type logbook struct {
	// nolog logbook and joiner
}

// SetLevel is part of mlog.Logbook interface implementation
func (*logbook) SetLevel(string, mlog.Level) error {
	// Logging is disabled, return error
	return mlog.ErrDisabledLogging
}

// Levels is part of mlog.Logbook interface implementation
func (*logbook) Levels() mlog.Levels {
	// Logging is disabled, no logging levels
	lvs := make(mlog.Levels)
	return lvs
}

// Joiner is part of mlog.Logbook interface implementation
func (lb *logbook) Joiner() mlog.Joiner {
	return lb
}

// Join is part of mlog.Joiner interface implementation
func (*logbook) Join(string) mlog.Logger {
	return &logger{}
}

// NewLogbook retruns logbook without no logging functionality
func NewLogbook() mlog.Logbook {
	return &logbook{}
}

// Interface implementation check
var (
	_ mlog.Logbook = &logbook{}
	_ mlog.Logger  = &logger{}
)
