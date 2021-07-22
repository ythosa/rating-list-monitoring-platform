package dto

import (
	"github.com/go-playground/validator/v10"
)

type UserCredentials struct {
	Nickname string `json:"nickname" validate:"required,min=4,max=10"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

func (d *UserCredentials) Validate(validate *validator.Validate) error {
	return validate.Struct(d)
}
