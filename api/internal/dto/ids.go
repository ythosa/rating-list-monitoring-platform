package dto

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type IDs struct {
	IDs []uint `json:"ids" validate:"required"`
}

func (d *IDs) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return fmt.Errorf("error while validating ids: %w", err)
	}

	return nil
}
