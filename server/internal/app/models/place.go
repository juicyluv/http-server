package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Place struct {
	Id   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}

type UpdatePlaceInput struct {
	Name *string `json:"name"`
}

func (p *Place) Validate() error {
	return validation.ValidateStruct(
		p,
		validation.Field(&p.Name, is.Alphanumeric, validation.Required),
	)
}
