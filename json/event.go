package json

import (
	"bytes"
	"runtime"
	"strconv"
	"time"

	"unicode/utf8"

	"go.melnyk.org/mlog"
)

type event struct {
	buffer bytes.Buffer
}

// Interface implementation check
var (
	_ mlog.Event = &event{}
)

func (evt *event) String(name string, value string) {
	evt.buffer.WriteByte('"')
	evt.buffer.WriteString(name)
	evt.buffer.WriteString(`":"`)
	evt.escapeString(value)
	evt.buffer.WriteString(`", `)
}

func (evt *event) Int(name string, value int) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendInt(b, (int64)(value), 10)
	evt.buffer.WriteByte('"')
	evt.buffer.WriteString(name)
	evt.buffer.WriteString(`":`)
	evt.buffer.Write(b)
	evt.buffer.WriteString(`, `)
}

func (evt *event) Uint(name string, value uint) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendUint(b, (uint64)(value), 10)
	evt.buffer.WriteByte('"')
	evt.buffer.WriteString(name)
	evt.buffer.WriteString(`":`)
	evt.buffer.Write(b)
	evt.buffer.WriteString(`, `)
}

func (evt *event) Hex(name string, value uint) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendUint(b, (uint64)(value), 16)
	evt.buffer.WriteByte('"')
	evt.buffer.WriteString(name)
	evt.buffer.WriteString(`":"0x`)
	evt.buffer.Write(b)
	evt.buffer.WriteString(`", `)
}

func (evt *event) Error(name string, value error) {
	evt.buffer.WriteByte('"')
	evt.buffer.WriteString(name)
	if value != nil {
		evt.buffer.WriteString(`":"`)
		evt.escapeString(value.Error())
		evt.buffer.WriteString(`", `)
	} else {
		evt.buffer.WriteString(`":null, `)
	}
}

// Formating part

func (evt *event) escapeString(s string) {
	start := 0
	for curr, c := range s {
		if c >= 0x20 && c != '\\' && c != '"' && c != '\u2028' && c != '\u2029' && c < utf8.RuneSelf {
			continue
		}
		if start < curr {
			evt.buffer.WriteString(s[start:curr])
		}
		start = curr + utf8.RuneLen(c)
		switch c {
		case '\\':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('\\')
		case '"':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('"')
		case '\n':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('n')
		case '\f':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('f')
		case '\b':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('b')
		case '\r':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('r')
		case '\t':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('t')
		case '\u2028', '\u2029':
			evt.buffer.WriteByte('\\')
			evt.buffer.WriteByte('n')
		default:
			// Don't encode non-ascii symbols
		}
	}
	if start < len(s) {
		evt.buffer.WriteString(s[start:])
	}
}

func (evt *event) tail() {
	var buf [64]byte
	b := buf[:0]
	ts := time.Now()
	b = ts.AppendFormat(b, timestampFormat)

	// copy timestamp
	evt.buffer.WriteString(`"ts":"`)
	evt.buffer.Write(b)
	evt.buffer.WriteString(`"}\n`)
}

func (evt *event) header(lv mlog.Level, name string) {

	representation := `{"level":"unknown", "logger:"`

	switch lv {
	case mlog.Fatal:
		representation = `{"level":"fatal", "logger:"`
	case mlog.Panic:
		representation = `{"level":"panic", "logger:"`
	case mlog.Error:
		representation = `{"level":"error", "logger:"`
	case mlog.Warning:
		representation = `{"level":"warning", "logger:"`
	case mlog.Info:
		representation = `{"level":"info", "logger:"`
	case mlog.Verbose:
		representation = `{"level":"verbose", "logger:"`
	}

	evt.buffer.WriteString(representation)
	evt.buffer.WriteString(name)
	evt.buffer.WriteString(`", `)
}

func (evt *event) frame(f *runtime.Frame) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendUint(b, (uint64)(f.Line), 10)
	evt.buffer.WriteByte('"')
	evt.buffer.WriteString(f.Function)
	evt.buffer.WriteByte('(')
	evt.buffer.WriteString(f.File)
	evt.buffer.WriteByte(':')
	evt.buffer.Write(b)
	evt.buffer.WriteString(`)"`)
}
