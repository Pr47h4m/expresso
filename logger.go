package expresso

import (
	"fmt"
	"net/url"
)

// Log levels constants used to indicate the severity of log messages.
const (
	LogLevelInfo  = "INFO"  // Informational messages.
	LogLevelError = "ERROR" // Error messages indicating something went wrong.
	LogLevelDebug = "DEBUG" // Debug messages for development and troubleshooting.
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
	l.logs = append(l.logs, Log{Level: LogLevelInfo, Message: message})
}

// Error adds an error log entry to the Logger.
func (l *Logger) Error(message string) {
	l.logs = append(l.logs, Log{Level: LogLevelError, Message: message})
}

// Debug adds a debug log entry to the Logger.
func (l *Logger) Debug(message string) {
	l.logs = append(l.logs, Log{Level: LogLevelDebug, Message: message})
}

// Dump prints all the logged messages to the console, formatted with colors
// based on their HTTP method and log level. The method and status code are
// displayed at the top, followed by each log entry.
func (l *Logger) Dump() {
	if len(l.logs) == 0 {
		return // Exit if there are no logs to print.
	}

	logStr := "+---\n" // Start log dump with a separator line.

	// Colorize and format the HTTP method.
	switch l.Method {
	case "GET":
		logStr += "\033[0;32m" + "| GET" + "\033[0m" + " "
	case "POST":
		logStr += "\033[0;33m" + "| POST" + "\033[0m" + " "
	case "PUT":
		logStr += "\033[0;34m" + "| PUT" + "\033[0m" + " "
	case "DELETE":
		logStr += "\033[0;35m" + "| DELETE" + "\033[0m" + " "
	case "PATCH":
		logStr += "\033[0;36m" + "| PATCH" + "\033[0m" + " "
	case "OPTIONS":
		logStr += "\033[0;37m" + "| OPTIONS" + "\033[0m" + " "
	case "HEAD":
		logStr += "\033[0;38m" + "| HEAD" + "\033[0m" + " "
	case "TRACE":
		logStr += "\033[0;39m" + "| TRACE" + "\033[0m" + " "
	case "CONNECT":
		logStr += "\033[0;40m" + "| CONNECT" + "\033[0m" + " "
	case "LINK":
		logStr += "\033[0;41m" + "| LINK" + "\033[0m" + " "
	case "UNLINK":
		logStr += "\033[0;42m" + "| UNLINK" + "\033[0m" + " "
	default:
		logStr += "\033[0;43m" + "| UNKNOWN" + "\033[0m" + " "
	}

	// Append the request path and status code, colorized based on the status range.
	logStr += fmt.Sprint(l.Path, " ")

	if l.StatusCode >= 200 && l.StatusCode < 300 {
		logStr += "\033[0;32m" + fmt.Sprint(l.StatusCode) + "\033[0m" + "\n"
	} else if l.StatusCode >= 300 {
		logStr += "\033[0;35m" + fmt.Sprint(l.StatusCode) + "\033[0m" + "\n"
	}

	// Print each log entry with color based on its level.
	for _, log := range l.logs {
		switch log.Level {
		case LogLevelInfo:
			logStr += "\033[0;32m" + "| " + log.Level + "\033[0m" + " "
		case LogLevelError:
			logStr += "\033[0;31m" + "| " + log.Level + "\033[0m" + " "
		case LogLevelDebug:
			logStr += "\033[0;34m" + "| " + log.Level + "\033[0m" + " "
		}
		logStr += log.Message + "\n"
	}
	logStr += "+---" // End log dump with a separator line.
	fmt.Println(logStr)
}
