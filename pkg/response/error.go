package response

import (
	"errors"
	"net/http"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Err     error  `json:"error"`
	Code    int    `json:"code"`
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
	InternalServerError = NewErrorResponse(http.StatusInternalServerError, "Internal Server Error")
)
