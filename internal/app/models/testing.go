package models

import "testing"

func TestUser(t *testing.T) *User {
	return &User{
		Email:             "user@mail.com",
		Username:          "admin",
		EncryptedPassword: "qwerty",
	}
}
