package utils

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

// HTTP util constants
const (
	contentTypeKey             = "Content-Type"
	acceptKey                  = "Accept"
	applicationJsonContentType = "application/json"
	textPlainContentType       = "text/plain"

	invalidProxyProtocolErrMsg = "Invalid proxy protocol '%s'. Valid values are %v"
	proxyUrlParseErrMsg        = "Unable to parse proxy URL : %v"
	proxyUsedMsg               = "Using proxy %s for making the request"
	createRequestErrMsg        = "Error creating base http request : %v"
	makeRequestErrMsg          = "Error making http request : %v"
	readResponseErrMsg         = "Error reading response body : %v"
)

var (
	// Skip TLS
	SkipTLS = false
	// ProxyEnabled enable proxy settings
	ProxyEnabled = false
	// ProxyProtocol protocol of the proxy server, http or https
	ProxyProtocol string
	// ProxyHost hostname of the proxy server
	ProxyHost string
	// ProxyPort port of the proxy server
	ProxyPort string
	// validProtocols
	validProtocols = []string{"http", "https"}
)

// InvalidProxyProtocolError represents an error when the proxy protocol is not valid
type InvalidProxyProtocolError string

// Error returns the formatted InvalidProxyProtocolError
func (ipp InvalidProxyProtocolError) Error() string {
	return fmt.Sprintf(invalidProxyProtocolErrMsg, string(ipp), validProtocols)
}

// ProxyUrlParseError represents an error when proxy url cannot be parsed
type ProxyUrlParseError struct {
	Err error
}

// Error returns the formatted ProxyUrlParseError
func (pup ProxyUrlParseError) Error() string {
	return fmt.Sprintf(proxyUrlParseErrMsg, pup.Err)
}

// CreateRequestError represents an error when creating a http request fails
type CreateRequestError struct {
	Err error
}

// Error returns teh formatted CreateRequestError
func (cr CreateRequestError) Error() string {
	return fmt.Sprintf(createRequestErrMsg, cr.Err)
}

// MakeRequestError represents an error when
type MakeRequestError struct {
	Err error
}

// Error returns teh formatted MakeRequestError
func (mr MakeRequestError) Error() string {
	return fmt.Sprintf(makeRequestErrMsg, mr.Err)
}

// ReadResponseError represents an error when the response cannot be read
type ReadResponseError struct {
	Err error
}

// Error returns teh formatted ReadResponseError
func (rr ReadResponseError) Error() string {
	return fmt.Sprintf(readResponseErrMsg, rr.Err)
}

// Auth represents authentication information
type Auth struct {
	Username string
	Password string
}

// RequestBody body represents the format of a request body
type RequestBody struct {
	Json []byte
	Text string
}

// Result represents the result of a http request
type Result struct {
	Body   []byte
	Status string
}

// Request represents an HTTP request
type Request struct {
	Url     string
	Method  string
	Auth    Auth
	Body    RequestBody
	Cnf     HTTPCnf
	Request *http.Request
	Result  Result
}

// ValidationResult represents the result of validation method
type ValidationResult struct {
	Valid    bool     `json:"valid"`
	Messages []string `json:"messages"`
}

// Response represents a generic http response message
type Response struct {
	Message string `json:"message"`
}

// ErrResponse represents a generic http error response message
type ErrResponse struct {
	Error string `json:"error"`
}

// GetRequesterIP gets the requester IP from the request headers and returns the IP
func GetRequesterIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusString(code int) string {
	return fmt.Sprintf("%d %s", code, http.StatusText(code))
}

// NewRequest creates a base http request based on the URL method and credentials provided in the Request struct
// The method write the created request back into the Request struct
// The method returns an error if the request creation fails
func (r *Request) NewRequest() error {

	var err error

	if r.Cnf.SkipTLS {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if r.Body.Json != nil {

		r.Request, err = http.NewRequest(r.Method, r.Url, bytes.NewBuffer(r.Body.Json))

		if err != nil {
			return CreateRequestError{Err: err}
		}

		r.Request.Header.Set(contentTypeKey, applicationJsonContentType)
		r.Request.Header.Set(acceptKey, applicationJsonContentType)

	} else if r.Body.Text != "" {

		r.Request, err = http.NewRequest(r.Method, r.Url, strings.NewReader(r.Body.Text))

		if err != nil {
			return CreateRequestError{Err: err}
		}

		r.Request.Header.Set(contentTypeKey, textPlainContentType)

	} else {

		r.Request, err = http.NewRequest(r.Method, r.Url, nil)

		if err != nil {
			return CreateRequestError{Err: err}
		}

	}

	if r.Auth.Username != "" && r.Auth.Password != "" {
		r.Request.SetBasicAuth(r.Auth.Username, r.Auth.Password)
	}

	return err
}

// HttpRequest makes an http request to a remote server
// The response body and the status of the http response is registered into the request struct
// The method returns an error if there is a problem with making the request or while
// reading the response from the remote server
func (r *Request) HttpRequest() error {

	client := &http.Client{}

	if r.Cnf.ProxyEnable == true {
		if err := r.Cnf.Validate(); err != nil {
			return err
		}

		proxyUrl, err := url.Parse(r.Cnf.GetProxyUrl())
		if err != nil {
			return ProxyUrlParseError{Err: err}
		}

		log := LogFormatter{Msg: fmt.Sprintf(proxyUsedMsg, proxyUrl)}
		log.Debug().Println(log.Out)

		transport := &http.Transport{
			Proxy: http.ProxyURL(proxyUrl),
		}

		client.Transport = transport
	}

	authorisation := r.Request.Header.Get("Authorization")
	r.Request.Header.Set("Authorization", "")
	log := LogFormatter{Msg: fmt.Sprintf("Request : %v", r.Request)}
	log.Debug().Println(log.Out)
	log = LogFormatter{Msg: fmt.Sprintf("Request Url : %v", r.Request.URL)}
	log.Debug().Println(log.Out)
	log = LogFormatter{Msg: fmt.Sprintf("Request Headers : %v", r.Request.Header)}
	log.Debug().Println(log.Out)
	log = LogFormatter{Msg: fmt.Sprintf("Request Body : %v", r.Request.Body)}
	log.Debug().Println(log.Out)
	r.Request.Header.Set("Authorization", authorisation)

	resp, err := client.Do(r.Request)
	if err != nil {
		return MakeRequestError{Err: err}
	}

	r.Result.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return ReadResponseError{Err: err}
	}

	r.Result.Status = resp.Status

	resp.Header.Set("Authorization", "")
	log = LogFormatter{Msg: fmt.Sprintf("Response : %v", resp)}
	log.Debug().Println(log.Out)
	log = LogFormatter{Msg: fmt.Sprintf("Response Headers : %v", resp.Header)}
	log.Debug().Println(log.Out)
	log = LogFormatter{Msg: fmt.Sprintf("Response Status : %v", resp.Status)}
	log.Debug().Println(log.Out)
	log = LogFormatter{Msg: fmt.Sprintf("Response Body : %v", string(r.Result.Body))}
	log.Debug().Println(log.Out)

	err = resp.Body.Close()
	if err != nil {
		return MakeRequestError{Err: err}
	}

	return err
}
