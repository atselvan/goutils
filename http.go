package utils

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net"
	"net/http"
)

const (
	SuccessStatusCode             = 200
	CreatedStatusCode             = 201
	BadRequestStatusCode          = 400
	UnauthorizedStatusCode        = 401
	NotFoundStatusCode            = 404
	MethodNotAllowedStatusCode    = 405
	InternalServerErrorStatusCode = 500

	SuccessStatus             = "200 OK"
	CreatedStatus             = "201 Created"
	BadReqStatus              = "400 Bad Request"
	UnauthorizedStatus        = "401 Unauthorized"
	NotFoundStatus            = "404 Not Found"
	MethodNotAllowed          = "405 Method Not Allowed"
	InternalServerErrorStatus = "500 Internal Server Error"

	JsonMarshalErrorStr   = "JSON Marshal Error"
	JsonUnmarshalErrorStr = "JSON Unmarshal Error"
	ApiAuthErrorStr       = "401 unauthorized. Please pass username and password to the API"
)

// Response represents a response message
type Response struct {
	Message string `json:"message"`
}

// ErrResponse represents an error message
type ErrResponse struct {
	Error string `json:"error"`
}

// getRequesterIP finds and returns the IP of the http requester
func getRequesterIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = r.Header.Get("X-Forwarded-For")
	}
	if ip == "" {
		ip, _, _ = net.SplitHostPort(r.RemoteAddr)
	}
	return ip
}

// BasicAuthCheck check if basic authentication credentials are provided while making a http request
// The method return an error if something goes wrong
func BasicAuthCheck(r *http.Request) error {
	var err error
	if r.Header.Get("Authorization") == "" {
		err = errors.New(ApiAuthErrorStr)
	}
	return err
}

// ReadRequestBody reads request body in json format, unmarshal's the json writes the data into the address
func ReadRequestBody(r *http.Request, out interface{}) error {

	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if len(reqBody) > 0 {
		err := json.Unmarshal(reqBody, out)
		if err != nil {
			return err
		}
	}

	return nil
}

// WriteHTTPResp writes an http response
func WriteHTTPResp(w http.ResponseWriter, r *http.Request, responseCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	out, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		WriteErrorResp(w, r, InternalServerErrorStatusCode, err)
	}
	w.WriteHeader(responseCode)
	_, err = w.Write(out)
	if err != nil {
		WriteErrorResp(w, r, InternalServerErrorStatusCode, err)
	}
}

// WriteInfoResp writes an http info response and logs the outcome
func WriteInfoResp(w http.ResponseWriter, r *http.Request, statusCode int, response string) {
	WriteHTTPResp(w, r, statusCode, Response{response})
	Logger{r, response}.Info()
}

// WriteWarnResp writes an http warn response and logs the outcome
func WriteWarnResp(w http.ResponseWriter, r *http.Request, statusCode int, response string) {
	WriteHTTPResp(w, r, statusCode, Response{response})
	Logger{r, response}.Warn()
}

// WriteErrResp writes an http error response and logs the outcome
func WriteErrorResp(w http.ResponseWriter, r *http.Request, statusCode int, err error) {
	response := ErrResponse{err.Error()}
	WriteHTTPResp(w, r, statusCode, response)
	Logger{r, err.Error()}.Error()
}
