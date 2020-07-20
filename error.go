package utils

import (
	"fmt"
)

const (
	missingMandatoryParamErrMsg = "Missing mandatory parameter(s) : %v"
	regexCompileErrMsg          = "Unable to compile regex : %v"
	jsonMarshalErrMsg           = "JSON marshal error : %v"
	jsonUnMarshalErrMsg         = "JSON unmarshal error : %v"
	yamlMarshalErrMsg           = "YAML marshal error : %v"
	yamlUnmarshalErrMsg         = "YAML unmarshal error : %v"
)

// MissingMandatoryParamError represents an error when a mandatory parameter is missing
type MissingMandatoryParamError []string

// Error returns the formatted MissingMandatoryParamError
func (mpe MissingMandatoryParamError) Error() string {
	return fmt.Sprintf(missingMandatoryParamErrMsg, []string(mpe))
}

// RegexCompileError represents an error when a regex compilation fails
type RegexCompileError struct {
	Err error
}

// Error returns the formatted RegexCompileError
func (rc RegexCompileError) Error() string {
	return fmt.Sprintf(regexCompileErrMsg, rc.Err)
}

// JSONMarshalError represents an error when json marshal fails
type JSONMarshalError struct {
	Err error
}

// Error returns the formatted JSONMarshalError
func (jm JSONMarshalError) Error() string {
	return fmt.Sprintf(jsonMarshalErrMsg, jm.Err)
}

// JSONUnMarshalError represents an error when json unmarshal fails
type JSONUnMarshalError struct {
	Err error
}

// Error returns the formatted JSONUnMarshalError
func (jum JSONUnMarshalError) Error() string {
	return fmt.Sprintf(jsonUnMarshalErrMsg, jum.Err)
}

// YAMLMarshalError represents an error when yaml marshal fails
type YAMLMarshalError struct {
	Err error
}

// Error returns the formatted YAMLMarshalError
func (ym YAMLMarshalError) Error() string {
	return fmt.Sprintf(yamlMarshalErrMsg, ym.Err)
}

// YAMLUnMarshalError represents an error when yaml unmarshal fails
type YAMLUnMarshalError struct {
	Err error
}

// Error returns the formatted YAMLUnMarshalError
func (yum YAMLUnMarshalError) Error() string {
	return fmt.Sprintf(yamlUnmarshalErrMsg, yum.Err)
}
