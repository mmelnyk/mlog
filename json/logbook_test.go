package json

import (
	"bytes"
	"testing"

	"go.melnyk.org/mlog"
)

func TestLogbookNew(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)

	if logbook == nil {
		t.Fatal("Logbook has not been returned")
	}

	joiner := logbook.Joiner()
	if joiner == nil {
		t.Fatal("Joiner has not been returned")
	}

	logger := joiner.Join("test")
	if logger == nil {
		t.Fatal("Logger has not been returned")
	}
}

func TestLogbookLogLevels(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)

	if logbook == nil {
		t.Fatal("Logbook has not been returned")
	}

	// Check default levels

	levels := logbook.Levels()

	if len(levels) != 1 {
		t.Fatal("Levels should have only DEFAULT entry instead of", levels)
	}

	if val, ok := levels[mlog.Default]; !ok || val != mlog.Fatal {
		t.Fatal("Level for DEFAULT should be FATAL instead of", levels)
	}

	// Check modified default levels

	if err := logbook.SetLevel(mlog.Default, mlog.Info); err != nil {
		t.Fatal("Error is not expected for SetLevel call")
	}

	levels = logbook.Levels()

	if len(levels) != 1 {
		t.Fatal("Levels should have only DEFAULT entry instead of", levels)
	}

	if val, ok := levels[mlog.Default]; !ok || val != mlog.Info {
		t.Fatal("Level for DEFAULT should be INFO instead of", levels)
	}

	// Check levels for test1 logger

	if err := logbook.SetLevel("test1", mlog.Error); err != nil {
		t.Fatal("Error is not expected for SetLevel call")
	}

	levels = logbook.Levels()

	if len(levels) != 2 {
		t.Fatal("Levels should have DEFAULT & test1 entries instead of", levels)
	}

	if val, ok := levels[mlog.Default]; !ok || val != mlog.Info {
		t.Fatal("Level for DEFAULT should be INFO instead of", levels)
	}

	if val, ok := levels["test1"]; !ok || val != mlog.Error {
		t.Fatal("Level for test1 should be ERROR instead of", levels)
	}

	// Check default for new logger test2

	_ = logbook.Joiner().Join("test2")

	levels = logbook.Levels()

	if len(levels) != 3 {
		t.Fatal("Levels should have DEFAULT & test1 & test2 entries instead of", levels)
	}

	if val, ok := levels[mlog.Default]; !ok || val != mlog.Info {
		t.Fatal("Level for DEFAULT should be INFO instead of", levels)
	}

	if val, ok := levels["test1"]; !ok || val != mlog.Error {
		t.Fatal("Level for test1 should be ERROR instead of", levels)
	}

	if val, ok := levels["test2"]; !ok || val != mlog.Info {
		t.Fatal("Level for test2 should be INFO instead of", levels)
	}

	// Default level change should touch only default dedicated loggers

	if err := logbook.SetLevel(mlog.Default, mlog.Warning); err != nil {
		t.Fatal("Error is not expected for SetLevel call")
	}

	levels = logbook.Levels()

	if len(levels) != 3 {
		t.Fatal("Levels should have DEFAULT & test1 & test2 entries instead of", levels)
	}

	if val, ok := levels[mlog.Default]; !ok || val != mlog.Warning {
		t.Fatal("Level for DEFAULT should be INFO instead of", levels)
	}

	if val, ok := levels["test1"]; !ok || val != mlog.Error {
		t.Fatal("Level for test1 should be ERROR instead of", levels)
	}

	if val, ok := levels["test2"]; !ok || val != mlog.Warning {
		t.Fatal("Level for test2 should be INFO instead of", levels)
	}

	// Loggers with default level should be converted to custom after SetLevel call

	if err := logbook.SetLevel("test2", mlog.Verbose); err != nil {
		t.Fatal("Error is not expected for SetLevel call")
	}

	if err := logbook.SetLevel(mlog.Default, mlog.Fatal); err != nil {
		t.Fatal("Error is not expected for SetLevel call")
	}

	levels = logbook.Levels()

	if len(levels) != 3 {
		t.Fatal("Levels should have DEFAULT & test1 & test2 entries instead of", levels)
	}

	if val, ok := levels[mlog.Default]; !ok || val != mlog.Fatal {
		t.Fatal("Level for DEFAULT should be FATAL instead of", levels)
	}

	if val, ok := levels["test1"]; !ok || val != mlog.Error {
		t.Fatal("Level for test1 should be ERROR instead of", levels)
	}

	if val, ok := levels["test2"]; !ok || val != mlog.Verbose {
		t.Fatal("Level for test2 should be VERBOSE instead of", levels)
	}
}

func TestLogbookJoin(t *testing.T) {
	buf := &bytes.Buffer{}
	logbook := NewLogbook(buf)

	if logbook == nil {
		t.Fatal("Logbook has not been returned")
	}

	joiner := logbook.Joiner()
	if joiner == nil {
		t.Fatal("Joiner has not been returned")
	}

	logger1 := joiner.Join("test")
	if logger1 == nil {
		t.Fatal("Logger has not been returned")
	}

	logger2 := joiner.Join("test")
	if logger2 == nil {
		t.Fatal("Logger has not been returned")
	}

	if logger1 != logger2 {
		t.Fatal("Expected same interface value for same logger")
	}
}
