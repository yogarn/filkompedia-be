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

	UserNotFound   = NewErrorResponse(http.StatusNotFound, "User not found")
	UserUnverified = NewErrorResponse(http.StatusForbidden, "User is not verified")

	InvalidToken       = NewErrorResponse(http.StatusUnauthorized, "Token invalid")
	InvalidOTP         = NewErrorResponse(http.StatusUnauthorized, "OTP invalid")
	ExpiredToken       = NewErrorResponse(http.StatusUnauthorized, "Expired token")
	InvalidCredentials = NewErrorResponse(http.StatusUnauthorized, "Invalid credentials")

	RoleUnauthorized = NewErrorResponse(http.StatusUnauthorized, "Insufficient role")
)
