package nolog

import (
	"testing"

	"go.melnyk.org/mlog"
)

func TestNolog(t *testing.T) {
	nolog := NewLogbook()

	if nolog == nil {
		t.Fatal("Logbook has not been returned")
	}

	joiner := nolog.Joiner()
	if joiner == nil {
		t.Fatal("Joiner has not been returned")
	}

	logger := joiner.Join("test")
	if logger == nil {
		t.Fatal("Logger has not been returned")
	}

	if err := nolog.SetLevel("test", mlog.Info); err != mlog.ErrDisabledLogging {
		t.Fatal("Expected ErrDisabledLogging instead of ", err)
	}

	levels := nolog.Levels()
	if len(levels) != 0 {
		t.Fatal("Levels should be empty instead of ", levels)
	}

	logger.Fatal("fatal")
	logger.Panic("panic")
	logger.Error("error")
	logger.Warning("warning")
	logger.Info("info")
	logger.Verbose("verbose")
	logger.Event(mlog.Verbose, nil)
}
