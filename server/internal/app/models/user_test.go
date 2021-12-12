package models_test

import (
	"testing"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_HashPassword(t *testing.T) {
	u := models.TestUser(t)
	assert.NoError(t, u.HashPassword())
	assert.NotEmpty(t, u.EncryptedPassword)
}

func TestUser_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		u       func() *models.User
		isValid bool
	}{
		{
			name: "default",
			u: func() *models.User {
				return models.TestUser(t)
			},
			isValid: true,
		},
		{
			name: "empty email",
			u: func() *models.User {
				return &models.User{
					Email:             "",
					Username:          "user",
					EncryptedPassword: "qwerty",
				}
			},
			isValid: false,
		},
		{
			name: "no password",
			u: func() *models.User {
				return &models.User{
					Email:             "user@mail.com",
					Username:          "user",
					EncryptedPassword: "",
				}
			},
			isValid: false,
		},
		{
			name: "invalid email",
			u: func() *models.User {
				return &models.User{
					Email:             "user",
					Username:          "user",
					EncryptedPassword: "qwerty",
				}
			},
			isValid: false,
		},
		{
			name: "invalid username",
			u: func() *models.User {
				return &models.User{
					Email:             "user@mail.com",
					Username:          "---/",
					EncryptedPassword: "qwerty123",
				}
			},
			isValid: false,
		},
		{
			name: "password less than 6",
			u: func() *models.User {
				return &models.User{
					Email:             "user@mail.com",
					Username:          "---/",
					EncryptedPassword: "asdf",
				}
			},
			isValid: false,
		},
		{
			name: "password more than 20",
			u: func() *models.User {
				return &models.User{
					Email:             "user@mail.com",
					Username:          "---/",
					EncryptedPassword: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaasdfweqf",
				}
			},
			isValid: false,
		},
		{
			name: "password is not alphabetic",
			u: func() *models.User {
				return &models.User{
					Email:             "user@mail.com",
					Username:          "---/",
					EncryptedPassword: "aaaaaa---//",
				}
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				assert.NoError(t, tc.u().Validate())
			} else {
				assert.Error(t, tc.u().Validate())
			}
		})
	}
}
