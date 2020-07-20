package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const (
	httpLogFormat         = "\t|%s %3d %s| %12s |%s %-6s %s| %s | %s"
	infoLogLevel          = "INFO"
	debugLogLevel         = "DEBUG"
	invalidLogLevelErrMsg = "Invalid log level '%s'. Valid values are %v"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	orange  = "\033[48;2;255;165;0m"
	reset   = "\033[0m"
)

var (
	validLogLevels = []string{infoLogLevel, debugLogLevel}
)

// InvalidLogLevelError represents an error when the log level is not valid
type InvalidLogLevelError string

// Error returns the formatted InvalidLogLevelError
func (ill InvalidLogLevelError) Error() string {
	return fmt.Sprintf(invalidLogLevelErrMsg, string(ill), validLogLevels)
}

// LogFormatter represents log message details
type LogFormatter struct {
	Request    *http.Request
	StatusCode int
	Msg        string
	ErrMsg     error
	Out        string
}

// StatusCodeColor is the ANSI color for appropriately logging http status code to a terminal.
func (l *LogFormatter) StatusCodeColor() string {
	code := l.StatusCode

	switch {
	case code >= http.StatusOK && code < http.StatusMultipleChoices:
		return green
	case code >= http.StatusMultipleChoices && code < http.StatusBadRequest:
		return yellow
	case code >= http.StatusBadRequest && code < http.StatusInternalServerError:
		return orange
	default:
		return red
	}
}

// MethodColor is the ANSI color for appropriately logging http method to a terminal.
func (l *LogFormatter) MethodColor() string {
	method := l.Request.Method

	switch method {
	case "GET":
		return blue
	case "POST":
		return cyan
	case "PUT":
		return yellow
	case "DELETE":
		return red
	default:
		return reset
	}
}

// ResetColor resets all escape attributes.
func (l *LogFormatter) ResetColor() string {
	return reset
}

// Logging represents the logging configuration details
type Logging struct {
	Level string `json:"level" yaml:"level"`
}

// GetMsg returns a formatted log message
func (l *LogFormatter) GetMsg() string {
	return strings.Replace(l.Msg, "\n", "", -1)
}

// GetErrMsg returns a formatted log message
func (l *LogFormatter) GetErrMsg() string {
	return strings.Replace(l.ErrMsg.Error(), "\n", "", -1)
}

// GetLogMsg formats a message based on the values set for Request and Message set for the Logger receiver
// If Request variable is nil only the message will be returned
// else a formatted string will the request details along with the message will be returned
func (l *LogFormatter) GetLogMsg() string {

	if l.Msg != "" && l.ErrMsg != nil {
		l.Out = l.GetMsg() + " : " + l.GetErrMsg()
	} else if l.Msg != "" {
		l.Out = l.GetMsg()
	} else {
		l.Out = l.GetErrMsg()
	}

	if l.Request != nil {
		statusColor := l.StatusCodeColor()
		methodColor := l.MethodColor()
		resetColor := l.ResetColor()

		l.Out = fmt.Sprintf(httpLogFormat,
			statusColor, l.StatusCode, resetColor,
			GetRequesterIP(l.Request),
			methodColor, l.Request.Method, resetColor,
			l.Request.RequestURI,
			l.Out)
	}

	return l.Out
}

// Info generates the output log message and returns a new Info logger
func (l *LogFormatter) Info() *log.Logger {
	l.GetLogMsg()
	return log.New(os.Stdout, "[INFO ] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Warn generates the output log message and returns a new Warn logger
func (l *LogFormatter) Warn() *log.Logger {
	l.GetLogMsg()
	return log.New(os.Stdout, "[WARN ] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Error generates the output log message and returns a new Error logger
func (l *LogFormatter) Error() *log.Logger {
	l.GetLogMsg()
	return log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Debug generates the output log message and returns a new Debug logger
func (l *LogFormatter) Debug() *log.Logger {
	l.GetLogMsg()
	if LogLevel == debugLogLevel {
		return log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return log.New(ioutil.Discard, "", log.LstdFlags)
}
