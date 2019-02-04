package clog

import "runtime"

/// Logger contract defines methods that must be available for a Logger.
///
/// Debug must write a message in DEBUG level.
///
/// Debugf must write a formatted message in DEBUG level.
///
/// Info must write a message in INFO level.
///
/// Infof must write a formatted message in INFO level.
///
/// Warn must write a message in WARN level.
///
/// Warnf must write a formatted message in WARN level.
///
/// Error must write an error, message that explaining the error and where its occurred in ERROR level.
/// To trace message, use Trace function and skip 1.
///
/// Errorf must write a formatted message and where its occurred in ERROR level.
/// To trace message, use Trace function and skip 1.
type Logger interface {
	Debug(msg string)
	Debugf(format string, args ...interface{})
	Info(msg string)
	Infof(format string, args ...interface{})
	Warn(msg string)
	Warnf(format string, args ...interface{})
	Error(msg string, err error)
	Errorf(format string, args ...interface{})
	Fatal(msg string, err error)
	Fatalf(format string, args ...interface{})
}

// Levels
const (
	LevelPanic = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
)

// Config
const (
	// Env Config Keys
	EnvLogLevel = "LOG_LEVEL"
	// Default Values
	DefaultLevel = LevelDebug
)

/// log is a singleton logger instance
var log Logger

/// Get retrieve singleton logger instance
func Get() Logger {
	if log == nil {
		panic("nbs-go/clog: no logger implementation has been registered")
	}
	return log
}

/// Register logger instance
func Register(l Logger) {
	if l == nil {
		panic("nbs-go/clog: logger to be registered is nil")
	}
	log = l
}

/// Trace retrieve where the code is being called and returns full path of file where the error occurred
func Trace(skip int) (string, int) {
	_, file, line, ok := runtime.Caller(skip + 1)
	if !ok {
		file = "<???>"
		line = 1
	}
	return file, line
}
