package mlog

import "errors"

const (
	// Default is name for loggers with default level
	Default = "DEFAULT"
)

var (
	// ErrDisabledLogging indicates when logging is disabled
	ErrDisabledLogging = errors.New("Logging is disabled")
)
