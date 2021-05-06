package api

import (
	"encoding/json"
	"net/http"
)

const DefaultErrorMessage = "An error has occurred. Please try again."

// ErroeMessage is a wrapper for a message to help with returning an error
// message along with an error status code
type ErrorMessage struct {
	Message string `json:"message"`
}

// newErrorMessage will return a new errorMessage given a message
func NewErrorMessage(msg string) *ErrorMessage {
	res := &ErrorMessage{
		Message: msg,
	}
	return res
}

func DefaultError(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(NewErrorMessage(message))
}
