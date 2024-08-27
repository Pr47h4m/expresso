package expresso

import (
	"fmt"
	"net/url"
)

const (
	LogLevelInfo  = "INFO"
	LogLevelError = "ERROR"
	LogLevelDebug = "DEBUG"
)

type Log struct {
	Level   string
	Message string
}

type Logger struct {
	Path   *url.URL
	Method string
	logs   []Log
}

func NewLogger(req *Request) *Logger {
	return &Logger{
		Path:   req.Path,
		Method: req.Method,
	}
}

func (l *Logger) Info(message string) {
	l.logs = append(l.logs, Log{Level: LogLevelInfo, Message: message})
}

func (l *Logger) Error(message string) {
	l.logs = append(l.logs, Log{Level: LogLevelError, Message: message})
}

func (l *Logger) Debug(message string) {
	l.logs = append(l.logs, Log{Level: LogLevelDebug, Message: message})
}

func (l *Logger) Dump() {
	if len(l.logs) == 0 {
		return
	}
	logStr := ""
	logStr += fmt.Sprintln(l.Method, l.Path)
	for _, log := range l.logs {
		logStr += fmt.Sprintln(log.Level, log.Message)
	}
	fmt.Println(logStr)
}
