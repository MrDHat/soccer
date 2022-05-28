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
	LoginInputEmailEmpty        = "LoginInputEmailEmpty"
	LoginInputPasswordEmpty     = "LoginInputPasswordEmpty"
	UserAlreadyExists           = "UserAlreadyExists"
	UserNotFound                = "UserNotFound"
	NoUserIdInContext           = "NoUserIdInContext"
	NoUserTokenInContext        = "NoUserTokenInContext"

	// team errors
	TeamNotFound = "TeamNotFound"

	// player errors
	PlayerNotFound = "PlayerNotFound"
)

var ErrorCode = map[string]string{
	// 500 errors
	InternalServerError: "100000",

	// 404 errors
	NotFound:       "400000",
	UserNotFound:   "400001",
	TeamNotFound:   "400002",
	PlayerNotFound: "400003",

	// 400 errors
	InvalidRequestData:          "200000",
	SignupInputNameEmpty:        "200001",
	SignupInputEmailEmpty:       "200002",
	SignupInputPasswordEmpty:    "200003",
	SignupInputPasswordTooShort: "200004",
	LoginInputEmailEmpty:        "200005",
	LoginInputPasswordEmpty:     "200007",
	UserAlreadyExists:           "200007",

	// 401 errors
	Unauthorized:         "300000",
	NoUserIdInContext:    "300001",
	NoUserTokenInContext: "300002",
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
	LoginInputEmailEmpty:        "The email field cannot be empty",
	LoginInputPasswordEmpty:     "The password field cannot be empty",
	UserAlreadyExists:           "A user with this email already exists",
	UserNotFound:                "The user could not be found",
	NoUserIdInContext:           "The user id could not be found in the context",
	NoUserTokenInContext:        "The user token could not be found in the context",

	TeamNotFound: "The team could not be found",

	PlayerNotFound: "The player could not be found",
}
