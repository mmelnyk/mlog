package testlog

import (
	"go.melnyk.org/mlog"
)

type event struct{}

func (*event) String(string, string) {}
func (*event) Int(string, int)       {}
func (*event) Uint(string, uint)     {}
func (*event) Hex(string, uint)      {}
func (*event) Error(string, error)   {}

type logger struct {
	// testlog logger
}

func (*logger) Verbose(string) {}
func (*logger) Info(string)    {}
func (*logger) Warning(string) {}
func (*logger) Error(string)   {}
func (*logger) Panic(string)   {}
func (*logger) Fatal(string)   {}

func (*logger) Event(_ mlog.Level, cb func(lg mlog.Event)) {
	if cb != nil {
		cb(&event{})
	}
}

type logbook struct {
	// testlog logbook and joiner
}

// SetLevel is part of mlog.Logbook interface implementation
func (*logbook) SetLevel(string, mlog.Level) error {
	// Logging is disabled, return error
	return mlog.ErrDisabledLogging
}

// Levels is part of mlog.Logbook interface implementation
func (*logbook) Levels() mlog.Levels {
	// Return dummy values
	lvs := make(mlog.Levels)
	lvs[mlog.Default] = mlog.Verbose
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
	_ mlog.Event   = &event{}
)
