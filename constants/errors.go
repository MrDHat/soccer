package constants

// Server error constants
const (
	// common errors
	InternalServerError = "InternalServerError"
	InvalidRequestData  = "InvalidRequestData"
	Unauthorized        = "Unauthorized"
	NotFound            = "NotFound"
	Conflict            = "Conflict"
)

var ErrorCode = map[string]string{
	InternalServerError: "000001",
	InvalidRequestData:  "000002",
	Unauthorized:        "000003",
	NotFound:            "000004",
	Conflict:            "000005",
}

// ErrorString returns the string version of the error which is sent to the user
var ErrorString = map[string]string{
	InternalServerError: "We're sorry! Looks like something went wrong",
	InvalidRequestData:  "The request failed because it contained an invalid value",
	Unauthorized:        "We're sorry! You are not authorized to perform this action",
	NotFound:            "The resource requested could not be found",
	Conflict:            "An item already exists with this name",
}
