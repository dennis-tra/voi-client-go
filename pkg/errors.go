package voi

import (
	"fmt"
	"net/http"
)

// ErrorVoiUnexpectedResponseCode is returned if the response code was not anticipated
type ErrorVoiUnexpectedResponseCode struct {
	// Response contains the actual http response from the endpoint
	Response *http.Response
}

// Error is a method to comply to the error interface
func (*ErrorVoiUnexpectedResponseCode) Error() string {
	return fmt.Sprintf("The response code is unhandled and therefore was unexpected")
}

// ErrorVoiBadRequest represents an error that occurs if the provided data was in the wrong format or in other means
// invalid
type ErrorVoiBadRequest struct {
	// Response contains the actual http response from the endpoint
	Response *http.Response
}

// Error is a method to comply to the error interface
func (*ErrorVoiBadRequest) Error() string {
	return fmt.Sprintf("The provided payload was not expected")
}

// ErrorUserAlreadySignedUp is returned if a sign up attempt was made for a user that was already registered
type ErrorUserAlreadySignedUp struct {
	Response *http.Response
}

// Error is a method to comply to the error interface
func (*ErrorUserAlreadySignedUp) Error() string {
	return fmt.Sprintf("The user is already registered")
}

// ErrorVoiUnauthroized ...
type ErrorVoiUnauthroized struct {
}

// Error is a method to comply to the error interface
func (*ErrorVoiUnauthroized) Error() string {
	return fmt.Sprintf("The provided payload was not expected")
}
