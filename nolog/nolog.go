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

func (lb *logbook) SetLevel(string, mlog.Level) error {
	return nil
}

func (lb *logbook) Levels() mlog.Levels {
	lvs := make(mlog.Levels)
	lvs[mlog.Default] = mlog.None
	return lvs
}

func (lb *logbook) Joiner() mlog.Joiner {
	return lb
}

func (lb *logbook) Join(string) mlog.Logger {
	return &logger{}
}

// NewLogbook retruns logbook without no logging functionality
func NewLogbook() mlog.Logbook {
	return &logbook{}
}
