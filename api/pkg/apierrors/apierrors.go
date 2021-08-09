package apierrors

import "errors"

type APIError struct {
	Message string `json:"message"`
}

func NewAPIError(err error) *APIError {
	return &APIError{Message: err.Error()}
}

func (e *APIError) Error() string {
	return e.Message
}

var (
	InvalidRefreshToken        = NewAPIError(errors.New("invalid token"))
	InvalidAuthorizationHeader = NewAPIError(errors.New("invalid authorization header"))
	InvalidQueryIDParam        = NewAPIError(errors.New("invalid query id param"))
)
