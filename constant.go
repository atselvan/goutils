package utils

const (
	SuccessStatusCode             = 200
	CreatedStatusCode             = 201
	FoundStatusCode				  = 302
	BadRequestStatusCode          = 400
	UnauthorizedStatusCode        = 401
	NotFoundStatusCode            = 404
	MethodNotAllowedStatusCode    = 405
	InternalServerErrorStatusCode = 500

	SuccessStatus             = "200 OK"
	CreatedStatus             = "201 Created"
	FoundStatus				  = "302 Found"
	BadReqStatus              = "400 Bad Request"
	UnauthorizedStatus        = "401 Unauthorized"
	NotFoundStatus            = "404 Not Found"
	MethodNotAllowed          = "405 Method Not Allowed"
	InternalServerErrorStatus = "500 Internal Server Error"

	JsonMarshalErrorStr   = "JSON Marshal Error"
	JsonUnmarshalErrorStr = "JSON Unmarshal Error"
	ApiAuthErrorStr       = "401 unauthorized. Please pass username and password to the API"
)
