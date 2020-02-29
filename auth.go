package utils

import (
	"encoding/base64"
	"net/http"
	"strings"
)

// BasicAuth represents basic authentication parameters
type BasicAuth struct {
	Username string
	Password string
}

// Check checks if basic authentication credentials are provided to the API while making a http request
// The method returns an error if the credentials are not provided
func (b *BasicAuth) Check(r *http.Request) error {
	if r.Header.Get("Authorization") == "" {
		return Error{Message: BasicAuthErrMsg}.NewError()
	}
	return nil
}

// Get checks if the Authorization header is set correctly in the request and will try to
// get the Basic Authorization credentials (username and password) from the header
// If the credentials are retrieved successfully then the method returns the username and password
// The method returns an error if the header is not set or if the decoding fails
func (b *BasicAuth) Get(r *http.Request) error {
	if err := b.Check(r); err != nil {
		return err
	}

	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(auth) != 2 || auth[0] != "Basic" {
		return Error{Message: BasicAuthErrMsg}.NewError()
	}

	dAuth, err := base64.StdEncoding.DecodeString(auth[1])
	if err != nil {
		return Error{Message: BasicAuthErrMsg}.NewError()
	}

	cred := strings.SplitN(string(dAuth), ":", 2)

	if len(cred) != 2 {
		return Error{Message: BasicAuthErrMsg}.NewError()
	}

	b.Username = cred[0]
	b.Password = cred[1]

	return nil
}
