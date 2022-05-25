package utils

// Response is the generic struct for sending a response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"json,omitempty"`
}
