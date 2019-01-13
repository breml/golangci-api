package logutil

type Func func(format string, args ...interface{})

type Log interface {
	Fatalf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Debugf(key string, format string, args ...interface{})

	Child(name string) Log
	SetLevel(level LogLevel)
}

type LogLevel int

const (
	// debug message, write to debug logs only by logutils.Debug
	LogLevelDebug LogLevel = 0

	// information messages, don't write too much messages,
	// only useful ones: they are shown when running with -v
	LogLevelInfo LogLevel = 1

	// hidden errors: non critical errors: work can be continued, no need to fail whole program;
	// tests will crash if any warning occurred.
	LogLevelWarn LogLevel = 2

	// only not hidden from user errors: whole program failing, usually
	// error logging happens in 1-2 places: in the "main" function.
	LogLevelError LogLevel = 3
)
