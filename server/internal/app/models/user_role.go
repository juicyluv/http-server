package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type UserRole struct {
	Id   uint   `json:"id"`
	Role string `json:"role"`
}

type UpdateUserRoleInput struct {
	Role *string `json:"role"`
}

func (u *UserRole) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Role, is.Alpha),
	)
}
