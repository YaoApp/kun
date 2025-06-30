package log

import (
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
)

// F the log fields
type F map[string]interface{}

// Logger the logrus logger
type Logger struct {
	*logrus.Logger
}

// Entry the log entry
type Entry struct {
	*logrus.Entry
}

// Level type
type Level uint32

// Format type
type Format uint32

const (
	// JSON JSON Format
	JSON Format = iota
	// TEXT Text Format
	TEXT = 1 // 1
)

// These are the different logging levels. You can set the logging level to log
// on your instance of logger, obtained with `logrus.New()`.
const (
	// PanicLevel level, highest level of severity. Logs and then calls panic with the
	// message passed to Debug, Info, ...
	PanicLevel Level = iota
	// FatalLevel level. Logs and then calls `logger.Exit(1)`. It will exit even if the
	// logging level is set to Panic.
	FatalLevel
	// ErrorLevel level. Logs. Used for errors that should definitely be noted.
	// Commonly used for hooks to send errors to an error tracking service.
	ErrorLevel
	// WarnLevel level. Non-critical entries that deserve eyes.
	WarnLevel
	// InfoLevel level. General operational entries about what's going on inside the
	// application.
	InfoLevel
	// DebugLevel level. Usually only enabled when debugging. Very verbose logging.
	DebugLevel
	// TraceLevel level. Designates finer-grained informational events than the Debug.
	TraceLevel
)

// StandardLogger returns the standard logger
func StandardLogger() *Logger {
	return &Logger{logrus.StandardLogger()}
}

// SetFormatter 0=JSON, 1=TEXT
func SetFormatter(format Format) {
	switch format {
	case JSON:
		logrus.SetFormatter(&logrus.JSONFormatter{})
	case TEXT:
		logrus.SetFormatter(&logrus.TextFormatter{})
	}
}

// SetOutput log writer
func SetOutput(out io.Writer) {
	logrus.SetOutput(out)
}

// SetLevel log level
func SetLevel(level Level) {
	logrus.SetLevel(logrus.Level(level))
}

// GetLevel get log level
func GetLevel() Level {
	return Level(logrus.GetLevel())
}

// With fields
func With(fields F) *Entry {
	return &Entry{logrus.WithFields(logrus.Fields(fields))}
}

// Trace trace log
func (entry *Entry) Trace(message string, v ...interface{}) {
	entry.Entry.Trace(fmt.Sprintf(message, v...))
}

// Trace trace log
func Trace(message string, v ...interface{}) {
	logrus.Trace(fmt.Sprintf(message, v...))
}

// Debug debug log
func (entry *Entry) Debug(message string, v ...interface{}) {
	entry.Entry.Debug(fmt.Sprintf(message, v...))
}

// Debug debug log
func Debug(message string, v ...interface{}) {
	logrus.Debug(fmt.Sprintf(message, v...))
}

// Info info log
func (entry *Entry) Info(message string, v ...interface{}) {
	entry.Entry.Info(fmt.Sprintf(message, v...))
}

// Info info log
func Info(message string, v ...interface{}) {
	logrus.Info(fmt.Sprintf(message, v...))
}

// Warn warn log
func (entry *Entry) Warn(message string, v ...interface{}) {
	entry.Entry.Warn(fmt.Sprintf(message, v...))
}

// Warn warn log
func Warn(message string, v ...interface{}) {
	logrus.Warn(fmt.Sprintf(message, v...))
}

// Error error log
func (entry *Entry) Error(message string, v ...interface{}) {
	entry.Entry.Error(fmt.Sprintf(message, v...))
}

// Error error log
func Error(message string, v ...interface{}) {
	logrus.Error(fmt.Sprintf(message, v...))
}

// Fatal fatal log Calls os.Exit(1) after logging
func (entry *Entry) Fatal(message string, v ...interface{}) {
	entry.Entry.Fatal(fmt.Sprintf(message, v...))
}

// Fatal fatal log Calls os.Exit(1) after logging
func Fatal(message string, v ...interface{}) {
	logrus.Fatal(fmt.Sprintf(message, v...))
}

// Panic panic log  Calls panic() after logging
func (entry *Entry) Panic(message string, v ...interface{}) {
	entry.Entry.Panic(fmt.Sprintf(message, v...))
}

// Panic panic log  Calls panic() after logging
func Panic(message string, v ...interface{}) {
	logrus.Panic(fmt.Sprintf(message, v...))
}
