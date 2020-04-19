package testlog

import (
	"testing"

	"go.melnyk.org/mlog"
)

func TestTestlog(t *testing.T) {
	testlog := NewLogbook()

	if testlog == nil {
		t.Fatal("Logbook has not been returned")
	}

	joiner := testlog.Joiner()
	if joiner == nil {
		t.Fatal("Joiner has not been returned")
	}

	logger := joiner.Join("test")
	if logger == nil {
		t.Fatal("Logger has not been returned")
	}

	if err := testlog.SetLevel("test", mlog.Info); err != mlog.ErrDisabledLogging {
		t.Fatal("Expected ErrDisabledLogging instead of ", err)
	}

	levels := testlog.Levels()
	if len(levels) != 1 {
		t.Fatal("Levels should have only DEFAULT logger instead of ", levels)
	}

	logger.Fatal("fatal")
	logger.Panic("panic")
	logger.Error("error")
	logger.Warning("warning")
	logger.Info("info")
	logger.Verbose("verbose")
	logger.Event(mlog.Verbose, func(e mlog.Event) {
		e.String("string", "value")
		e.Int("int", 1000)
		e.Uint("uint", 2000)
		e.Hex("hex", 3000)
		e.Error("error", nil)
	})
}
