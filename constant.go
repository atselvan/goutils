package utils

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





	fileNotFoundErrMsg              = "File not found"
	fileNotFoundErrDetail           = "the file %s was not found"
	fileReadErrMsg                  = "Unable to read file"
	fileCreateErrMsg                = "Unable to create a new file"
	fileOpenErrMsg                  = "Unable to open file"
	fileWriteErrMsg                 = "Unable to write to the file"
	JsonMarshalErrMsg               = "JSON Marshal Error"
	JsonUnmarshalErrMsg             = "JSON Unmarshal Error"
	YamlMarshalErrMsg               = "YAML Marshal Error"
	YamlUnmarshalErrMsg             = "YAML Unmarshal Error"
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
	promptConfirmErrMsg             = "input can be either y/n"
	promptSelectMoreMsg             = "Do you want to select one more ?"
)

var (
	// Default log level
	LogLevel      = "INFO"
	ProxyEnabled  = false
	ProxyProtocol string
	ProxyHost     string
	ProxyPort     string
)
