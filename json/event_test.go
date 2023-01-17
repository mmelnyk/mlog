package json

import (
	"errors"
	"testing"
)

// TODO
func TestEventString(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.String("name", "value")

	if ev.buffer.String() != `"name":"value", ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventEscapingString(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.String("name", "\"value\ttab\nnew\r\b\falso\"\u2028\u2029 \\ \x03\x02the end")

	if ev.buffer.String() != `"name":"\"value\ttab\nnew\r\b\falso\"\n\n \\ the end", ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventIntPos(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.Int("name", 1234567890)

	if ev.buffer.String() != `"name":1234567890, ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventIntNeg(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.Int("name", -1234567890)

	if ev.buffer.String() != `"name":-1234567890, ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventUIntPos(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.Uint("name", 1234567890)

	if ev.buffer.String() != `"name":1234567890, ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventUIntNeg(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	var v uint
	v = v - 1234567890
	ev.Uint("name", v)

	if ev.buffer.String() != `"name":18446744072474983726, ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventHex(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.Hex("name", 1234567890)

	if ev.buffer.String() != `"name":"0x499602d2", ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventHexNeg(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	var v uint
	v = v - 1234567890
	ev.Hex("name", v)

	if ev.buffer.String() != `"name":"0xffffffffb669fd2e", ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventErrorNil(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.Error("name", nil)

	if ev.buffer.String() != `"name":null, ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}

func TestEventError(t *testing.T) {
	ev := getEvent()

	if ev == nil {
		t.Fatal("Expected not nil event")
	}

	ev.Error("name", errors.New("Test error"))

	if ev.buffer.String() != `"name":"Test error", ` {
		t.Fatal("ev.String build not expected output:", ev.buffer.String())
	}

	putEvent(ev)
}
