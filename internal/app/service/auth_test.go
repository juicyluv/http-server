package service_test

import (
	"testing"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
	"github.com/ellywynn/http-server/internal/app/service"
	"github.com/stretchr/testify/assert"
)

func TestAuthService_Login(t *testing.T) {
	s, teardown := service.NewTestService(t, dbURL)
	defer teardown("users")

	email := "user@gmail.com"
	password := "qwerty123"

	testCases := []struct {
		name    string
		login   func() interfaces.AuthLoginStruct
		isValid bool
	}{
		{
			name: "correct login",
			login: func() interfaces.AuthLoginStruct {
				return interfaces.AuthLoginStruct{
					Email:    email,
					Password: password,
				}
			},
			isValid: true,
		},
		{
			name: "incorrect email",
			login: func() interfaces.AuthLoginStruct {
				return interfaces.AuthLoginStruct{
					Email:    "testemail@",
					Password: password,
				}
			},
			isValid: false,
		},
		{
			name: "incorrect password",
			login: func() interfaces.AuthLoginStruct {
				return interfaces.AuthLoginStruct{
					Email:    email,
					Password: "qwerty12",
				}
			},
			isValid: false,
		},
		{
			name: "incorrect email and password",
			login: func() interfaces.AuthLoginStruct {
				return interfaces.AuthLoginStruct{
					Email:    "user@eee.com",
					Password: "qwert222y12",
				}
			},
			isValid: false,
		},
		{
			name: "empty email",
			login: func() interfaces.AuthLoginStruct {
				return interfaces.AuthLoginStruct{
					Email:    "",
					Password: password,
				}
			},
			isValid: false,
		},
		{
			name: "empty password",
			login: func() interfaces.AuthLoginStruct {
				return interfaces.AuthLoginStruct{
					Email:    email,
					Password: "",
				}
			},
			isValid: false,
		},
	}

	userId, err := s.User.SignUp(&models.User{
		Email:             email,
		Username:          "user",
		EncryptedPassword: password,
	})

	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)

	for _, tc := range testCases {
		if tc.isValid {
			t.Run(tc.name, func(t *testing.T) {
				if tc.isValid {
					u, err := s.Auth.LogIn(tc.login())
					assert.NoError(t, err)
					assert.NotNil(t, u)
				} else {
					u, err := s.Auth.LogIn(tc.login())
					assert.Error(t, err)
					assert.Nil(t, u)
				}
			})
		}
	}
}
