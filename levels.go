package mlog

// Level type for logging level
type Level uint32

// Predefined logging levels
const (
	// None is intended to turn off logging.
	None Level = iota
	// Fatal - Designates very severe error events that will presumably lead
	// the application to abort (non-recoverable event).
	Fatal
	// Error - Designates error events that might still allow
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
	case None:
		name = "NONE"
	case Fatal:
		name = "FATAL"
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
