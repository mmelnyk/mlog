package console

import (
	"bytes"
	"strings"
	"testing"

	"go.melnyk.org/mlog"
)

func utilFillLog(log mlog.Logger) {
	log.Verbose("verbose msg")
	log.Info("info msg")
	log.Warning("warning msg")
	log.Error("error msg")
	log.Panic("panic msg")
	log.Fatal("fatal msg")
}

func utilFillLogCB(log mlog.Logger) {
	log.Event(mlog.Verbose, func(e mlog.Event) {
		e.String("msg", "verbose msg")
	})
	log.Event(mlog.Info, func(e mlog.Event) {
		e.String("msg", "info msg")
	})
	log.Event(mlog.Warning, func(e mlog.Event) {
		e.String("msg", "warning msg")
	})
	log.Event(mlog.Error, func(e mlog.Event) {
		e.String("msg", "error msg")
	})
	log.Event(mlog.Panic, func(e mlog.Event) {
		e.String("msg", "panic msg")
	})
	log.Event(mlog.Fatal, func(e mlog.Event) {
		e.String("msg", "fatal msg")
	})
}

func utilTestLevel(t *testing.T, out string, level mlog.Level) {
	if inc := strings.Contains(out, "FATAL"); inc != (level >= mlog.Fatal) {
		t.Fatal("Failed FATAL level check", out)
	}

	if inc := strings.Contains(out, "fatal msg"); inc != (level >= mlog.Fatal) {
		t.Fatal("Failed FATAL level check (msg output)")
	}

	if inc := strings.Contains(out, "PANIC"); inc != (level >= mlog.Panic) {
		t.Fatal("Failed PANIC level check", out)
	}

	if inc := strings.Contains(out, "panic msg"); inc != (level >= mlog.Panic) {
		t.Fatal("Failed PANIC level check (msg output)")
	}

	if inc := strings.Contains(out, "ERROR"); inc != (level >= mlog.Error) {
		t.Fatal("Failed ERROR level check", out)
	}

	if inc := strings.Contains(out, "error msg"); inc != (level >= mlog.Error) {
		t.Fatal("Failed ERROR level check (msg output)")
	}

	if inc := strings.Contains(out, "WARNING"); inc != (level >= mlog.Warning) {
		t.Fatal("Failed WARNING level check")
	}

	if inc := strings.Contains(out, "warning msg"); inc != (level >= mlog.Warning) {
		t.Fatal("Failed WARNING level check (msg output)")
	}

	if inc := strings.Contains(out, "INFO"); inc != (level >= mlog.Info) {
		t.Fatal("Failed INFO level check")
	}

	if inc := strings.Contains(out, "info msg"); inc != (level >= mlog.Info) {
		t.Fatal("Failed INFO level check (msg output)")
	}

	if inc := strings.Contains(out, "VERBOSE"); inc != (level >= mlog.Verbose) {
		t.Fatal("Failed VERBOSE level check")
	}

	if inc := strings.Contains(out, "verbose msg"); inc != (level >= mlog.Verbose) {
		t.Fatal("Failed VERBOSE level check (msg output)")
	}
}

func TestLoggerFatal(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)
	logger := logbook.Joiner().Join("test")

	logbook.SetLevel("test", mlog.Fatal)
	utilFillLog(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Fatal)

	buf.Reset()
	utilFillLogCB(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Fatal)
}

func TestLoggerPanic(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)
	logger := logbook.Joiner().Join("test")

	logbook.SetLevel("test", mlog.Panic)
	utilFillLog(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Panic)

	buf.Reset()
	utilFillLogCB(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Panic)
}

func TestLoggerError(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)
	logger := logbook.Joiner().Join("test")

	logbook.SetLevel("test", mlog.Error)
	utilFillLog(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Error)

	buf.Reset()
	utilFillLogCB(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Error)
}

func TestLoggerWarning(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)
	logger := logbook.Joiner().Join("test")

	logbook.SetLevel("test", mlog.Warning)
	utilFillLog(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Warning)

	buf.Reset()
	utilFillLogCB(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Warning)
}

func TestLoggerInfo(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)
	logger := logbook.Joiner().Join("test")

	logbook.SetLevel("test", mlog.Info)
	utilFillLog(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Info)

	buf.Reset()
	utilFillLogCB(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Info)
}

func TestLoggerVerbose(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)
	logger := logbook.Joiner().Join("test")

	logbook.SetLevel("test", mlog.Verbose)
	utilFillLog(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Verbose)

	buf.Reset()
	utilFillLogCB(logger)
	utilTestLevel(t, string(buf.Bytes()), mlog.Verbose)
}
