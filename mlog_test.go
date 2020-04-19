package mlog

import "testing"

// This package defines interfaces only, so no unittest required

func TestBasicLevels(t *testing.T) {
	if Fatal.String() != "FATAL" {
		t.Fail()
	}

	if Error.String() != "ERROR" {
		t.Fail()
	}

	if Warning.String() != "WARNING" {
		t.Fail()
	}

	if Info.String() != "INFO" {
		t.Fail()
	}

	if Verbose.String() != "VERBOSE" {
		t.Fail()
	}

	var l Level = Verbose + 1
	if l.String() != "UNKNOWN" {
		t.Fail()
	}
}
