package models

import "testing"

// Returns default user entity
func TestUser(t *testing.T) *User {
	return &User{
		Email:             "user@mail.com",
		Username:          "admin",
		EncryptedPassword: "qwerty",
	}
}
