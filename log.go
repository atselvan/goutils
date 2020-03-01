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
	logFormat     = "%s"
	httpLogFormat = "%s | %s | %s | %d | %s"
)

// Logger represents the log message
type Logger struct {
	Request *http.Request
	Status  int
	Message string
	Err     error
	Out     string
}

// Logging represents the logging configuration details
type Logging struct {
	Level string `json:"level" yaml:"level"`
}

// GetLogLevel checks if the LOG_LEVEL env variable is set
// If the environment variable has the value "DEBUG" the method returns "DEBUG"
// else the function always returns the value "INFO"
func (l *Logger) GetLogLevel() string {
	switch LogLevel {
	case "DEBUG":
		return "DEBUG"
	default:
		return "INFO"
	}
}

// GetMessage returns a formatted log message
func (l *Logger) GetMessage() string {
	return strings.Replace(l.Message, "\n", "", -1)
}

// GetError returns a formatted error log
func (l *Logger) GetError() string {
	return strings.Replace(l.Err.Error(), "\n", "", -1)
}

// GetLogMessage formats a message based on the values set for Request and Message set for the Logger receiver
// If Request variable is nil only the message will be returned
// else a formatted string will the request details along with the message will be returned
func (l *Logger) GetLogMessage() string {
	if l.Request != nil {
		l.Out = fmt.Sprintf(httpLogFormat, GetRequesterIP(l.Request), l.Request.Method, l.Request.RequestURI, l.Status, l.GetMessage())
	} else if l.Message != "" && l.Err != nil {
		l.Out = fmt.Sprintf(logFormat, fmt.Sprintf("%v : %s", l.GetMessage(), l.GetError()))
	} else if l.Message != "" {
		l.Out = fmt.Sprintf(logFormat, l.GetMessage())
	} else {
		fmt.Println(l.Err)
		l.Out = fmt.Sprintf(logFormat, l.GetError())
	}
	return l.Out
}

// Info generates the output log message and returns a new Info logger
func (l *Logger) Info() *log.Logger {
	l.GetLogMessage()
	return log.New(os.Stdout, "INFO  : ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Warn generates the output log message and returns a new Warn logger
func (l *Logger) Warn() *log.Logger {
	l.GetLogMessage()
	return log.New(os.Stdout, "WARN  : ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Error generates the output log message and returns a new Error logger
func (l *Logger) Error() *log.Logger {
	l.GetLogMessage()
	return log.New(os.Stdout, "ERROR : ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Debug generates the output log message and returns a new Debug logger
func (l *Logger) Debug() *log.Logger {
	l.GetLogMessage()
	if l.GetLogLevel() == "DEBUG" {
		return log.New(os.Stdout, "DEBUG : ", log.Ldate|log.Ltime|log.Lshortfile)
	}
	return log.New(ioutil.Discard, "", log.LstdFlags)
}
