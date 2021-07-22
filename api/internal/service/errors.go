package service

import "net/http"

type Error struct {
	HTTPCode int `json:"-"`
	Message  string
}

func NewServiceError(HTTPCode int, message string) *Error {
	return &Error{
		HTTPCode: HTTPCode,
		Message:  message,
	}
}

func (s *Error) Error() string {
	return s.Message
}

func NewStatusBadRequest(err string) *Error {
	return NewServiceError(http.StatusBadRequest, err)
}

func NewInternalServerError(err string) *Error {
	return NewServiceError(http.StatusInternalServerError, err)
}
