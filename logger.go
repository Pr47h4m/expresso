package expresso

import (
	"fmt"
	"net/url"
)

// Log levels constants used to indicate the severity of log messages.
const (
	LogLevelInfo  = "INFO"
	LogLevelError = "ERROR"
	LogLevelDebug = "DEBUG"
)

// Log represents a single log entry with a severity level and a message.
type Log struct {
	Level   string // The severity level of the log (e.g., INFO, ERROR, DEBUG).
	Message string // The content of the log message.
}

// Logger captures and manages log entries associated with an HTTP request.
type Logger struct {
	Path       *url.URL // The URL path of the request.
	Method     string   // The HTTP method used for the request (e.g., GET, POST).
	logs       []Log    // A slice of Log entries recorded during the request.
	StatusCode int      // The HTTP status code that will be returned with the response.
}

// NewLogger creates and returns a new Logger instance, initializing it with the request's path and method.
func NewLogger(req *Request) *Logger {
	return &Logger{
		Path:       req.Path,
		Method:     req.Method,
		StatusCode: 200, // Default status code is set to 200 (OK).
	}
}

// Info adds an informational log entry to the Logger.
func (l *Logger) Info(message string) {
	l.addLog(LogLevelInfo, message)
}

// Error adds an error log entry to the Logger.
func (l *Logger) Error(message string) {
	l.addLog(LogLevelError, message)
}

// Debug adds a debug log entry to the Logger.
func (l *Logger) Debug(message string) {
	l.addLog(LogLevelDebug, message)
}

// addLog is a helper function to append a log entry to the logs slice.
func (l *Logger) addLog(level, message string) {
	l.logs = append(l.logs, Log{Level: level, Message: message})
}

// Dump prints all the logged messages to the console, formatted with colors
// based on their HTTP method and log level. The method and status code are
// displayed at the top, followed by each log entry.
func (l *Logger) Dump() {
	if len(l.logs) == 0 {
		return // Exit if there are no logs to print.
	}

	logStr := "+---\n" // Start log dump with a separator line.

	// Append the request method, path, and status code.
	logStr += fmt.Sprintf("%s %s %s\n", l.colorizeMethod(), l.Path, l.colorizeStatusCode())

	// Append each log entry with color based on its level.
	for _, log := range l.logs {
		logStr += fmt.Sprintf("%s %s\n", l.colorizeLevel(log.Level), log.Message)
	}

	logStr += "+---" // End log dump with a separator line.
	fmt.Println(logStr)
}

// colorizeMethod returns the colored string based on the HTTP method.
func (l *Logger) colorizeMethod() string {
	colorMap := map[string]string{
		"GET":     "\033[0;32mGET\033[0m",
		"POST":    "\033[0;33mPOST\033[0m",
		"PUT":     "\033[0;34mPUT\033[0m",
		"DELETE":  "\033[0;35mDELETE\033[0m",
		"PATCH":   "\033[0;36mPATCH\033[0m",
		"OPTIONS": "\033[0;37mOPTIONS\033[0m",
		"HEAD":    "\033[0;38mHEAD\033[0m",
		"TRACE":   "\033[0;39mTRACE\033[0m",
		"CONNECT": "\033[0;40mCONNECT\033[0m",
		"LINK":    "\033[0;41mLINK\033[0m",
		"UNLINK":  "\033[0;42mUNLINK\033[0m",
	}
	if color, found := colorMap[l.Method]; found {
		return "|" + color
	}
	return "| \033[0;43mUNKNOWN\033[0m"
}

// colorizeStatusCode returns the colored string based on the HTTP status code.
func (l *Logger) colorizeStatusCode() string {
	switch {
	case l.StatusCode >= 200 && l.StatusCode < 300:
		return fmt.Sprintf("\033[0;32m%d\033[0m", l.StatusCode)
	case l.StatusCode >= 300:
		return fmt.Sprintf("\033[0;35m%d\033[0m", l.StatusCode)
	default:
		return fmt.Sprintf("%d", l.StatusCode)
	}
}

// colorizeLevel returns the colored string based on the log level.
func (l *Logger) colorizeLevel(level string) string {
	colorMap := map[string]string{
		LogLevelInfo:  "\033[0;32m| INFO\033[0m",
		LogLevelError: "\033[0;31m| ERROR\033[0m",
		LogLevelDebug: "\033[0;34m| DEBUG\033[0m",
	}
	return colorMap[level]
}
