package dto

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type UserCredentials struct {
	Username string `json:"username" validate:"required,min=4,max=10"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

func (d *UserCredentials) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf("error while validating user credentials: %w", err)
	}

	return nil
}
