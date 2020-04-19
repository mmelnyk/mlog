package mlog

import (
	"bytes"
	"errors"
)

// Level type for logging level
type Level uint32

// Predefined logging levels
const (
	// Fatal - Designates very severe error events that will presumably lead
	// the application to abort (non-recoverable event).
	Fatal Level = iota
	// Panic - Designates severe error events that might still allow
	// the application to continue running (recoverable event).
	Panic
	// Error - Designates error events that might allow
	// the application to continue running.
	Error
	// Warning - Designates potentially harmful situations.
	Warning
	// Info - Designates informational messages that highlight the progress of
	// the application at coarse-grained level.
	Info
	// Verbose - Designates fine-grained informational events that are most
	// useful to debug an application.
	Verbose
)

func (l Level) String() (name string) {
	switch l {
	case Fatal:
		name = "fatal"
	case Panic:
		name = "panic"
	case Error:
		name = "error"
	case Warning:
		name = "warning"
	case Info:
		name = "info"
	case Verbose:
		name = "verbose"
	default:
		name = "unknown"
	}
	return
}

// UppercaseString returns an uppercase representation of the log level.
func (l Level) UppercaseString() (name string) {
	switch l {
	case Fatal:
		name = "FATAL"
	case Panic:
		name = "PANIC"
	case Error:
		name = "ERROR"
	case Warning:
		name = "WARNING"
	case Info:
		name = "INFO"
	case Verbose:
		name = "VERBOSE"
	default:
		name = "UNKNOWN"
	}
	return
}

// MarshalText marshals the Level to text.
func (l Level) MarshalText() ([]byte, error) {
	return []byte(l.String()), nil
}

// UnmarshalText unmarshals text to a level.
//
// In particular, this makes it easy to configure logging levels using YAML or JSON files.
func (l *Level) UnmarshalText(text []byte) error {
	if l == nil {
		return ErrUnmarshalNil
	}

	switch string(bytes.ToLower(text)) {
	case "fatal", "": // default value
		*l = Fatal
	case "panic":
		*l = Panic
	case "error":
		*l = Error
	case "warning":
		*l = Warning
	case "info":
		*l = Info
	case "verbose":
		*l = Verbose
	default:
		return errors.New("Unrecognized level value: " + string(bytes.ToLower(text)))
	}

	return nil
}
