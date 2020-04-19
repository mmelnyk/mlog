package mlog

// Levels represent levels of existed loggers
type Levels map[string]Level

// Logbook interface provides an access to Logbook implementation for an app
type Logbook interface {
	SetLevel(string, Level) error // Set level to logger; possible error - ErrDisabledLogging
	Levels() Levels               // Get levels of all loggers
	Joiner() Joiner               // Get joiner interface for logbook
}

// Joiner interface allows component to join Logbook via creating/getting named logger
type Joiner interface {
	Join(string) Logger // Join logger to logbook (in other words - get logger interface)
}

// Event provides an access to add custom fields/data to a log
type Event interface {
	String(string, string)
	Int(string, int)
	Uint(string, uint)
	Hex(string, uint)
	Error(string, error)
}

// Logger allows to add event/messages to logbook
type Logger interface {
	Verbose(string)
	Info(string)
	Warning(string)
	Error(string)
	Panic(string)
	Fatal(string)
	Event(Level, func(Event))
}
