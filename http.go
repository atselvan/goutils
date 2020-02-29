package utils

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
)

const (
	GetMethod    = "GET"
	PostMethod   = "POST"
	PutMethod    = "PUT"
	DeleteMethod = "DELETE"

	SuccessStatusCode             = 200
	CreatedStatusCode             = 201
	FoundStatusCode               = 302
	BadRequestStatusCode          = 400
	UnauthorizedStatusCode        = 401
	NotFoundStatusCode            = 404
	MethodNotAllowedStatusCode    = 405
	ConflictStatusCode            = 409
	InternalServerErrorStatusCode = 500

	SuccessStatus             = "200 OK"
	CreatedStatus             = "201 Created"
	FoundStatus               = "302 Found"
	BadReqStatus              = "400 Bad Request"
	UnauthorizedStatus        = "401 Unauthorized"
	NotFoundStatus            = "404 Not Found"
	MethodNotAllowedStatus    = "405 Method Not Allowed"
	ConflictStatus            = "409 Conflict"
	InternalServerErrorStatus = "500 Internal Server Error"
	SuccessMsg                = "Success OK"
	PathNotFound              = "Request path '%s' not found"

	WriteRespErrMsg                 = "Unable to write any response on the writer"
	BasicAuthErrMsg                 = "401 unauthorized: Basic authentication is required"
	InvalidProxyErrMsg              = "Invalid proxy details"
	InvalidProxyErrDetail           = "When proxy is enabled, proxy host and proxy port should be provided"
	InvalidProxyProtocolErrMsg      = "Invalid proxy protocol"
	InvalidProxyProtocolErrDetail   = "Proxy protocol should be either http or https"
	ProxyUrlParseErrMsg             = "Unable to parse proxy URL"
	proxyUsedMsg                    = "Using proxy %s for making the request"
	reqCreateErrMsg                 = "Error creating base request"
	httpReqErrMsg                   = "Error making HTTP request"
	httpReqReadErrMsg               = "Error reading request body"
	httpRespReadErrMsg              = "Error reading response body"
)

var (
	// Default log level
	LogLevel      = "INFO"
	ProxyEnabled  = false
	ProxyProtocol string
	ProxyHost     string
	ProxyPort     string
)

// Request represents an HTTP request
type Request struct {
	Url     string
	Method  string
	Auth    Auth
	Body    RequestBody
	HTTP    HTTP
	Proxy   Proxy
	Request *http.Request
	Result  struct {
		Body   []byte
		Status string
	}
}

// Request body represents the format of a request body
type RequestBody struct {
	Json []byte
	Text string
}

// Auth represents authentication information
type Auth struct {
	Username string
	Password string
}

// Proxy represents proxy details
type Proxy struct {
	Enable   bool   `json:"enable" yaml:"enable"`
	Protocol string `json:"protocol" yaml:"protocol"`
	Host     string `json:"host" yaml:"host"`
	Port     string `json:"port" yaml:"port"`
}

// HTTP represents HTTP configuration details
type HTTP struct {
	SkipTLS bool `json:"skip_tls" yaml:"skip_tls"`
}

// Result represents the result of a http request
type Result struct {
	Body   []byte
	Status string
}

// ValidationResult represents the result of validation method
type ValidationResult struct {
	Valid    bool     `json:"valid"`
	Messages []string `json:"messages"`
}

// Response represents an HTTP response message
type Response struct {
	Message string `json:"message"`
}

// ErrResponse represents an error response
type ErrResponse struct {
	Error Error `json:"error"`
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

// ReadRequestBody reads request body in json format, unmarshal's the json and writes the data into the out variable
// out variable should be a pointer to a structure into which the values are un-marshaled
func ReadRequestBody(r *http.Request, out interface{}) error {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return Error{Message: httpReqReadErrMsg, Detail: err.Error()}.NewError()
	}

	if len(reqBody) > 0 {
		err := json.Unmarshal(reqBody, out)
		if err != nil {
			return Error{Message: JsonUnmarshalErrMsg, Detail: err.Error()}.NewError()
		}
	}

	return nil
}

// NewRequest creates a base http request based on the URL method and credentials provided in the Request struct
// The method write the created request back into the Request struct
// The method returns an error if the request creation fails
func (r *Request) NewRequest() error {

	var err error

	if r.HTTP.SkipTLS {
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	if r.Body.Json != nil {

		r.Request, err = http.NewRequest(r.Method, r.Url, bytes.NewBuffer(r.Body.Json))

		if err != nil {
			return Error{Message: reqCreateErrMsg, Detail: err.Error()}.NewError()
		}

		r.Request.Header.Set("Content-Type", "application/json")
		r.Request.Header.Set("Accept", "application/json")

	} else if r.Body.Text != "" {

		r.Request, err = http.NewRequest(r.Method, r.Url, strings.NewReader(r.Body.Text))

		if err != nil {
			return Error{Message: reqCreateErrMsg, Detail: err.Error()}.NewError()
		}

		r.Request.Header.Set("Content-Type", "text/plain")

	} else {

		r.Request, err = http.NewRequest(r.Method, r.Url, nil)

		if err != nil {
			return Error{Message: reqCreateErrMsg, Detail: err.Error()}.NewError()
		}

	}

	r.Request.SetBasicAuth(r.Auth.Username, r.Auth.Password)

	return err
}

// ValidateProxyDetails check if required proxy details are provided
// The method returns an error if proxy details are not valid
func (r *Request) validateProxyDetails() error {
	if strings.TrimSpace(r.Proxy.Protocol) == "" {
		r.Proxy.Protocol = "https"
	}
	if r.Proxy.Protocol != "http" && r.Proxy.Protocol != "https" {
		return Error{Message: InvalidProxyProtocolErrMsg, Detail: InvalidProxyProtocolErrDetail}.NewError()
	}
	if strings.TrimSpace(r.Proxy.Host) == "" || strings.TrimSpace(r.Proxy.Port) == "" {
		return Error{Message: InvalidProxyErrMsg, Detail: InvalidProxyErrDetail}.NewError()
	}
	return nil
}

func (r *Request) getProxyUrl() string {
	return fmt.Sprintf("%s://%s:%s", r.Proxy.Protocol, r.Proxy.Host, r.Proxy.Port)
}

// HttpRequest makes an http request to a remote server
// The response body and the status of the http response is registered into the request struct
// The method returns an error if there is a problem with making the request or while
// reading the response from the remote server
func (r *Request) HttpRequest() error {

	client := &http.Client{}

	if r.Proxy.Enable == true {
		if err := r.validateProxyDetails(); err != nil {
			return err
		} else {
			proxyUrl, err := url.Parse(r.getProxyUrl())

			if err != nil {
				return Error{Message: ProxyUrlParseErrMsg, Detail: err.Error()}.NewError()
			}

			log := Logger{Message: fmt.Sprintf(proxyUsedMsg, proxyUrl)}
			log.Debug().Println(log.Out)

			transport := &http.Transport{
				Proxy: http.ProxyURL(proxyUrl),
			}

			client.Transport = transport
		}
	}

	authorisation := r.Request.Header.Get("Authorization")
	r.Request.Header.Set("Authorization", "")
	log := Logger{Message: fmt.Sprintf("Request : %v", r.Request)}
	log.Debug().Println(log.Out)
	log = Logger{Message: fmt.Sprintf("Request Url : %v", r.Request.URL)}
	log.Debug().Println(log.Out)
	log = Logger{Message: fmt.Sprintf("Request Headers : %v", r.Request.Header)}
	log.Debug().Println(log.Out)
	log = Logger{Message: fmt.Sprintf("Request Body : %v", r.Request.Body)}
	log.Debug().Println(log.Out)
	r.Request.Header.Set("Authorization", authorisation)

	resp, err := client.Do(r.Request)
	if err != nil {
		return Error{Message: httpReqErrMsg, Detail: err.Error()}.NewError()
	}

	r.Result.Body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return Error{Message: httpRespReadErrMsg, Detail: err.Error()}.NewError()
	}

	r.Result.Status = resp.Status

	resp.Header.Set("Authorization", "")
	log = Logger{Message: fmt.Sprintf("Response : %v", resp)}
	log.Debug().Println(log.Out)
	log = Logger{Message: fmt.Sprintf("Response Headers : %v", resp.Header)}
	log.Debug().Println(log.Out)
	log = Logger{Message: fmt.Sprintf("Response Status : %v", resp.Status)}
	log.Debug().Println(log.Out)
	log = Logger{Message: fmt.Sprintf("Response Body : %v", string(r.Result.Body))}
	log.Debug().Println(log.Out)

	err = resp.Body.Close()
	if err != nil {
		return Error{Message: httpReqErrMsg, Detail: err.Error()}.NewError()
	}

	return err
}

// WriteHTTPResp writes an http response onto the response writer
func WriteHTTPResp(w http.ResponseWriter, r *http.Request, responseCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")

	out, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log := Logger{Message: JsonMarshalErrMsg, Err: err}
		log.Error().Println(log.Out)

		w.WriteHeader(InternalServerErrorStatusCode)
		_, err = w.Write(out)
	}

	w.WriteHeader(responseCode)
	_, err = w.Write(out)

	if err != nil {
		l := Logger{Message: WriteRespErrMsg, Err: err}
		l.Error().Println(l.Out)

		w.WriteHeader(InternalServerErrorStatusCode)
		_, err = w.Write(out)
	}
}

// WriteInfoResp writes information onto the response writer
func WriteInfoResp(w http.ResponseWriter, r *http.Request, responseCode int, message string) {
	WriteHTTPResp(w, r, responseCode, Response{Message: message})
}

// WriteErrResp writes error response onto the response writer
func WriteErrResp(w http.ResponseWriter, r *http.Request, responseCode int, err error) {
	var errResponse Error
	errorInfo := strings.SplitN(err.Error(), ":", 2)
	if len(errorInfo) == 2 {
		errResponse = Error{
			Message: strings.TrimSpace(errorInfo[0]),
			Detail:  strings.TrimSpace(errorInfo[1]),
		}
	} else {
		errResponse = Error{Message: errorInfo[0]}
	}
	WriteHTTPResp(w, r, responseCode, ErrResponse{errResponse})
}
