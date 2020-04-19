package nolog

import (
	"go.melnyk.org/mlog"
)

type logger struct {
	// nolog logger
}

func (l *logger) Verbose(string) {}
func (l *logger) Info(string)    {}
func (l *logger) Warning(string) {}
func (l *logger) Error(string)   {}
func (l *logger) Fatal(string)   {}

func (l *logger) Event(level mlog.Level, cb func(lg mlog.Event)) {}

type logbook struct {
	// nolog logbook and joiner
}

// SetLevel is part of mlog.Logbook interface implementation
func (lb *logbook) SetLevel(string, mlog.Level) error {
	// Logging is disabled, return error
	return mlog.ErrDisabledLogging
}

// Levels is part of mlog.Logbook interface implementation
func (lb *logbook) Levels() mlog.Levels {
	// Logging is disabled, no logging levels
	lvs := make(mlog.Levels)
	return lvs
}

// Joiner is part of mlog.Logbook interface implementation
func (lb *logbook) Joiner() mlog.Joiner {
	return lb
}

// Join is part of mlog.Joiner interface implementation
func (lb *logbook) Join(string) mlog.Logger {
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
