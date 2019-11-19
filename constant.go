package utils

const (
	SuccessStatusCode             = 200
	CreatedStatusCode             = 201
	FoundStatusCode               = 302
	BadRequestStatusCode          = 400
	UnauthorizedStatusCode        = 401
	NotFoundStatusCode            = 404
	MethodNotAllowedStatusCode    = 405
	InternalServerErrorStatusCode = 500

	SuccessStatus             = "200 OK"
	CreatedStatus             = "201 Created"
	FoundStatus               = "302 Found"
	BadReqStatus              = "400 Bad Request"
	UnauthorizedStatus        = "401 Unauthorized"
	NotFoundStatus            = "404 Not Found"
	MethodNotAllowed          = "405 Method Not Allowed"
	InternalServerErrorStatus = "500 Internal Server Error"

	JsonMarshalErrorStr   = "JSON Marshal Error"
	JsonUnmarshalErrorStr = "JSON Unmarshal Error"
	ApiAuthErrorStr       = "401 unauthorized. Please pass username and password to the API"

	dateFormatRegex              = `([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))`
	dateFormatLayout             = "2006-01-02"
	regexCompileErr              = "REGEX_COMPILE_ERROR"
	regexCompileErrStr           = "there was an error compiling the regex"
	dateFormatErr                = "INVALID_DATE_FORMAT"
	dateFormatErrStr             = "date should be of the format YYYY-MM-DD"
	invalidDateErr               = "INVALID_DATE"
	invalidDateErrStr            = "there was an error while parsing the date : %v"
	greaterThanCurrentDateErrStr = "date should not be greater than current date"
	invalidYearErr               = "INVALID_YEAR"
	invalidYearErrStr            = "year should be between 1990 and %d"
)
