package response

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Err  error `json:"error"`
	Code int   `json:"code"`
}

func (e *ErrorResponse) Error() string {
	return e.Err.Error()
}

func NewErrorResponse(code int, message string) ErrorResponse {
	return ErrorResponse{
		Code: code,
		Err:  errors.New(message),
	}
}

var (
	InternalServerError = NewErrorResponse(http.StatusInternalServerError, "Internal server error")

	UserNotFound = NewErrorResponse(http.StatusNotFound, "User not found")

	InvalidToken       = NewErrorResponse(http.StatusUnauthorized, "Token invalid")
	InvalidCredentials = NewErrorResponse(http.StatusUnauthorized, "Invalid credentials")
)
