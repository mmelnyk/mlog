package mlog

import (
	"encoding/json"
	"testing"
)

// This package defines interfaces only, so no unittest required

func TestBasicLevels(t *testing.T) {
	if Fatal.String() != "fatal" {
		t.Fail()
	}

	if Panic.String() != "panic" {
		t.Fail()
	}

	if Error.String() != "error" {
		t.Fail()
	}

	if Warning.String() != "warning" {
		t.Fail()
	}

	if Info.String() != "info" {
		t.Fail()
	}

	if Verbose.String() != "verbose" {
		t.Fail()
	}

	var l Level = Verbose + 1
	if l.String() != "unknown" {
		t.Fail()
	}
}

func TestBasicUppercaseLevels(t *testing.T) {
	if Fatal.UppercaseString() != "FATAL" {
		t.Fail()
	}

	if Panic.UppercaseString() != "PANIC" {
		t.Fail()
	}

	if Error.UppercaseString() != "ERROR" {
		t.Fail()
	}

	if Warning.UppercaseString() != "WARNING" {
		t.Fail()
	}

	if Info.UppercaseString() != "INFO" {
		t.Fail()
	}

	if Verbose.UppercaseString() != "VERBOSE" {
		t.Fail()
	}

	var l Level = Verbose + 1
	if l.UppercaseString() != "UNKNOWN" {
		t.Fail()
	}
}

func TestBasicLevelMarshal(t *testing.T) {
	tests := []struct {
		L Level
	}{
		{Fatal},
		{Panic},
		{Error},
		{Warning},
		{Info},
		{Verbose},
	}

	expected := `[{"L":"fatal"},{"L":"panic"},{"L":"error"},{"L":"warning"},{"L":"info"},{"L":"verbose"}]`

	b, err := json.Marshal(tests)
	if err != nil {
		t.Fatalf("Marshal failed with error: %s", err.Error())
	}

	if string(b) != expected {
		t.Fatalf("Expected marshaled text does not match result: %s", string(b))
	}
}

func TestBasicLevelUnmarshal(t *testing.T) {
	obj := []struct {
		L Level
	}{}

	tests := `[{"L":"fatal"},{"L":"panic"},{"L":"error"},{"L":"warning"},{"L":"info"},{"L":"verbose"}]`
	expected := []struct {
		L Level
	}{
		{Fatal},
		{Panic},
		{Error},
		{Warning},
		{Info},
		{Verbose},
	}

	err := json.Unmarshal([]byte(tests), &obj)
	if err != nil {
		t.Fatalf("Unmarshal failed with error: %s", err.Error())
	}

	for k, v := range expected {
		if obj[k] != v {
			t.Errorf("Expected %s, but got %s", v, obj[k])
		}
	}
}

func TestBasicLevelUnmarshalErrors(t *testing.T) {
	var l1 Level
	var l2 *Level

	err := l1.UnmarshalText([]byte(`errors`))
	if err == nil {
		t.Fatalf("Expected error on Unmarshal call, but got nil")
	}

	err = l2.UnmarshalText([]byte(`errors`))
	if err != ErrUnmarshalNil {
		t.Fatalf("Expected ErrUnmarshalNil error on Unmarshal call, but got %q", err)
	}
}
