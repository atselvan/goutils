package utils

import (
	"errors"
	"fmt"
)

// Error represents an error
type Error struct {
	Message string
	Detail  string
}

// NewError formats an error based on the provided message or detail or both in the error struct and returns an error
func (e Error) NewError() error {
	if e.Detail == "" {
		return errors.New(e.Message)
	} else {
		return errors.New(fmt.Sprintf("%s : %s", e.Message, e.Detail))
	}
}
