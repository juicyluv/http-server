package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int
	Email             string
	Username          string
	EncryptedPassword string
}

func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required, is.Alphanumeric, validation.Length(3, 20)),
		validation.Field(&u.EncryptedPassword, validation.Required, is.Alphanumeric, validation.Length(6, 20)),
	)
}

func (u *User) HashPassword() error {
	enc, err := encryptString(u.EncryptedPassword)
	if err != nil {
		return err
	}

	u.EncryptedPassword = enc
	return nil
}

func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
