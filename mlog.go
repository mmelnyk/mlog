package mlog

// Levels represent levels of existed loggers
type Levels map[string]Level

// Logbook interface provides an access to Logbook implementation for an app
type Logbook interface {
	SetLevel(string, Level) error
	Levels() Levels
	Joiner() Joiner
}

// Joiner interface allows component to join Logbook via creating/getting named logger
type Joiner interface {
	Join(string) Logger
}

// Event provides an access to add custom fields/data to a log
type Event interface {
	String(string, string)
	Int(string, int)
	Uint(string, uint)
	Hex(string, uint)
}

// Logger allows to add event/messages to logbook
type Logger interface {
	Verbose(string)
	Info(string)
	Warning(string)
	Error(string)
	Fatal(string)
	Event(Level, func(Event))
}
