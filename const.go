package mlog

import "errors"

const (
	// Default is name for loggers with default level
	Default = "DEFAULT"
)

var (
	// ErrDisabledLogging indicates when logging is disabled
	ErrDisabledLogging = errors.New("Logging is disabled")
	// ErrUnmarshalNil indicates error when Unmarshal called for nil pointer
	ErrUnmarshalNil = errors.New("Unmarshal to nil is not possible")
)
