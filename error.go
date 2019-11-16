package utils

import (
	"errors"
	"fmt"
	"strings"
)

// Error represents the error format
type Error struct {
	ErrStr string
	ErrMsg string
}

// getErrStr returns the error string
func (e Error) getErrStr() string {
	return strings.TrimSpace(e.ErrStr)
}

// getErrorMsg returns the error message
func (e Error) getErrMsg() string {
	return e.ErrMsg
}

// NewError formats and returns a error
func (e Error) NewError() error {
	if e.getErrStr() == "" {
		return errors.New(e.getErrMsg())
	} else {
		return errors.New(fmt.Sprintf("%s : %s", e.getErrStr(), e.getErrMsg()))
	}
}

// NewError creates and returns a new error and returns it
func NewError(errStr string) error {
	return errors.New(errStr)
}
