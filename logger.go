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
	Path       *url.URL
	Method     string
	logs       []Log
	StatusCode int
}

func NewLogger(req *Request) *Logger {
	return &Logger{
		Path:       req.Path,
		Method:     req.Method,
		StatusCode: 200,
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
	logStr := "+---\n"

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

	logStr += fmt.Sprint(l.Path, " ")

	if l.StatusCode >= 200 && l.StatusCode < 300 {
		logStr += "\033[0;32m" + fmt.Sprint(l.StatusCode) + "\033[0m" + "\n"
	} else if l.StatusCode >= 300 {
		logStr += "\033[0;35m" + fmt.Sprint(l.StatusCode) + "\033[0m" + "\n"
	}

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
	logStr += "+---"
	fmt.Println(logStr)
}
