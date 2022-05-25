package constants

import "fmt"

// Server error constants
const (
	// common errors
	InternalServerError = "InternalServerError"
	InvalidRequestData  = "InvalidRequestData"
	Unauthorized        = "Unauthorized"
	NotFound            = "NotFound"
	Conflict            = "Conflict"

	// user errors
	SignupInputNameEmpty        = "SignupInputNameEmpty"
	SignupInputEmailEmpty       = "SignupInputEmailEmpty"
	SignupInputPasswordEmpty    = "SignupInputPasswordEmpty"
	SignupInputPasswordTooShort = "SignupInputPasswordTooShort"
)

var ErrorCode = map[string]string{
	InternalServerError: "100000",
	InvalidRequestData:  "200000",
	Unauthorized:        "300000",
	NotFound:            "400000",
	Conflict:            "500000",

	SignupInputNameEmpty:        "200001",
	SignupInputEmailEmpty:       "200002",
	SignupInputPasswordEmpty:    "200003",
	SignupInputPasswordTooShort: "200004",
}

// ErrorString returns the string version of the error which is sent to the user
var ErrorString = map[string]string{
	InternalServerError: "We're sorry! Looks like something went wrong",
	InvalidRequestData:  "The request failed because it contained an invalid value",
	Unauthorized:        "We're sorry! You are not authorized to perform this action",
	NotFound:            "The resource requested could not be found",
	Conflict:            "An item already exists with this name",

	SignupInputNameEmpty:        "The name field cannot be empty",
	SignupInputEmailEmpty:       "The email field cannot be empty",
	SignupInputPasswordEmpty:    "The password field cannot be empty",
	SignupInputPasswordTooShort: fmt.Sprintf("The password is too short. It should be minimum %v characters", MinPasswordLength),
}
