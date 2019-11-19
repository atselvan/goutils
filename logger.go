package utils

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
)

const (
	logFormat     = "| %s"
	httpLogFormat = "| %s | %s | %s | %s"
)

// Logger represents the output log message
type Logger struct {
	Request *http.Request
	Message interface{}
}

// EnableDebug checks if the DEBUG_LOGS env variable is set to true
func (l Logger) EnableDebug() bool {
	if os.Getenv("DEBUG_LOGS") == "true" {
		return true
	} else {
		return false
	}
}

// Info writes information logs
func (l Logger) Info() {
	var out string
	infoLog := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)
	if l.Request != nil {
		out = fmt.Sprintf(httpLogFormat, getRequesterIP(l.Request), l.Request.Method, l.Request.RequestURI, l.Message)
	} else {
		out = fmt.Sprintf(logFormat, l.Message)
	}
	infoLog.Println(out)
}

// Warn writes warning logs
func (l Logger) Warn() {
	var out string
	warnLog := log.New(os.Stdout, "WARN: ", log.Ldate|log.Ltime)
	if l.Request != nil {
		out = fmt.Sprintf(httpLogFormat, getRequesterIP(l.Request), l.Request.Method, l.Request.RequestURI, l.Message)
	} else {
		out = fmt.Sprintf(logFormat, l.Message)
	}
	warnLog.Println(out)
}

// Error writes error logs
func (l Logger) Error() {
	var out string
	errLog := log.New(os.Stdout, "ERROR: ", log.Ldate|log.Ltime)
	if l.Request != nil {
		out = fmt.Sprintf(httpLogFormat, getRequesterIP(l.Request), l.Request.Method, l.Request.RequestURI, l.Message)
	} else {
		out = fmt.Sprintf(logFormat, l.Message)
	}
	errLog.Println(out)
	if l.EnableDebug() {
		debug.PrintStack()
	}
}
