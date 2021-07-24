package dto

import (
	"github.com/go-playground/validator/v10"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/pkg/validation"
	"gopkg.in/errgo.v2/fmt/errors"
)

type SigningUp struct {
	Username   string `json:"username" validate:"required,min=4,max=10"`
	Password   string `json:"password" validate:"required,min=5,max=20"`
	FirstName  string `json:"first_name" validate:"required,alpha,min=3"`
	MiddleName string `json:"middle_name" validate:"required,alpha,min=3"`
	LastName   string `json:"last_name" validate:"required,alpha,min=3"`
	Snils      string `json:"snils" validate:"required,numeric,len=11"`
}

func (d *SigningUp) Validate(validate *validator.Validate) error {
	if err := validate.Struct(d); err != nil {
		return errors.Newf("failed to validate dto: %s", err)
	}

	if err := validation.Snils(d.Snils); err != nil {
		return errors.Newf("invalid snils: %s", err)
	}

	return nil
}
