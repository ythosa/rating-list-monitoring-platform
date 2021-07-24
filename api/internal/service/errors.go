package service

import (
	"gopkg.in/errgo.v2/fmt/errors"
	"net/http"
)

type Error struct {
	HTTPCode int `json:"-"`
	Message  string
}

func NewError(HTTPCode int, err error) *Error {
	return &Error{
		HTTPCode: HTTPCode,
		Message:  err.Error(),
	}
}

func (s *Error) Error() string {
	return s.Message
}

func NewStatusBadRequest(err error) *Error {
	return NewError(http.StatusBadRequest, err)
}

func NewInternalServerError(err error) *Error {
	return NewError(http.StatusInternalServerError, err)
}

var (
	UserAlreadyExistsError         = NewError(http.StatusConflict, errors.New("user already exists"))
	InvalidUserNameOrPasswordError = NewError(http.StatusUnauthorized, errors.New("invalid username or password"))
)
