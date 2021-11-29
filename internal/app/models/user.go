package models

import (
	"errors"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id                int    `json:"id"`
	Email             string `json:"email"`
	Username          string `json:"username"`
	EncryptedPassword string `json:"password,omitempty"`
}

type UserUpdateInput struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

// Validates creating user struct
func (u *User) Validate() error {
	return validation.ValidateStruct(
		u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(&u.Username, validation.Required, is.Alphanumeric, validation.Length(3, 20)),
		validation.Field(&u.EncryptedPassword, validation.Required, is.Alphanumeric, validation.Length(6, 20)),
	)
}

// Hashes user password
func (u *User) HashPassword() error {
	enc, err := encryptString(u.EncryptedPassword)
	if err != nil {
		return err
	}

	u.EncryptedPassword = enc
	return nil
}

// Compare hashed user password with appropriate password
func (u *User) ComparePassword(password string) bool {
	return bcrypt.CompareHashAndPassword([]byte(u.EncryptedPassword), []byte(password)) == nil
}

// Returns encrypted password as string
func encryptString(password string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}

func (u UserUpdateInput) Validate() error {
	if u.Username == "" && u.Email == "" && u.Password == "" {
		return errors.New("object cannot be empty")
	}

	return nil
}
