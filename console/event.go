package console

import (
	"bytes"
	"runtime"
	"strconv"
	"time"
)

type event struct {
	buffer bytes.Buffer
}

func (evt *event) String(name string, value string) {
	evt.buffer.WriteString(name)
	evt.buffer.WriteByte('=')
	evt.buffer.WriteString(value)
	evt.buffer.WriteByte(' ')
}

func (evt *event) Int(name string, value int) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendInt(b, (int64)(value), 10)
	evt.buffer.WriteString(name)
	evt.buffer.WriteRune('=')
	evt.buffer.Write(b)
	evt.buffer.WriteRune(' ')
}

func (evt *event) Uint(name string, value uint) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendUint(b, (uint64)(value), 10)
	evt.buffer.WriteString(name)
	evt.buffer.WriteByte('=')
	evt.buffer.Write(b)
	evt.buffer.WriteByte(' ')
}

func (evt *event) Hex(name string, value uint) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendUint(b, (uint64)(value), 16)
	evt.buffer.WriteString(name)
	evt.buffer.WriteString("=0x")
	evt.buffer.Write(b)
	evt.buffer.WriteByte(' ')
}

// Formating part

func (evt *event) addTimestamp() {
	var buf [64]byte
	b := buf[:0]
	ts := time.Now()
	b = ts.AppendFormat(b, timestampFormat)
	evt.buffer.Write(b)
	// 3 spaces after
	for i := len(timestampFormat) - len(b) + 3; i > 0; i-- {
		evt.buffer.WriteByte(' ')
	}
}

func (evt *event) justJoinString(value string) {
	evt.buffer.WriteString(value)
	evt.buffer.WriteByte('\t')
}

func (evt *event) singleframe(f *runtime.Frame) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendUint(b, (uint64)(f.Line), 10)
	evt.buffer.WriteString(fieldCode)
	evt.buffer.WriteByte('=')
	evt.buffer.WriteString(f.Function)
	evt.buffer.WriteByte('(')
	evt.buffer.WriteString(f.File)
	evt.buffer.WriteByte(':')
	evt.buffer.Write(b)
	evt.buffer.WriteString(") ")
}

func (evt *event) frame(num int, f *runtime.Frame) {
	var buf [32]byte
	b := buf[:0]
	b = strconv.AppendUint(b, (uint64)(num), 10)
	evt.buffer.WriteString("\033[31m\t#")
	evt.buffer.Write(b)
	evt.buffer.WriteString("\t")
	evt.buffer.WriteString(f.Function)
	b = buf[:0]
	b = strconv.AppendUint(b, (uint64)(f.Line), 10)
	evt.buffer.WriteByte('(')
	evt.buffer.WriteString(f.File)
	evt.buffer.WriteByte(':')
	evt.buffer.Write(b)
	evt.buffer.WriteString(")\033[0m\n")
}
