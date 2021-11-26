package service_test

import (
	"testing"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/repository"
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
		login   func() *repository.LoginStruct
		isValid bool
	}{
		{
			name: "correct login",
			login: func() *repository.LoginStruct {
				return &repository.LoginStruct{
					Email:    email,
					Password: password,
				}
			},
			isValid: true,
		},
		{
			name: "incorrect email",
			login: func() *repository.LoginStruct {
				return &repository.LoginStruct{
					Email:    "testemail@",
					Password: password,
				}
			},
			isValid: false,
		},
		{
			name: "incorrect password",
			login: func() *repository.LoginStruct {
				return &repository.LoginStruct{
					Email:    email,
					Password: "qwerty12",
				}
			},
			isValid: false,
		},
		{
			name: "incorrect email and password",
			login: func() *repository.LoginStruct {
				return &repository.LoginStruct{
					Email:    "user@eee.com",
					Password: "qwert222y12",
				}
			},
			isValid: false,
		},
		{
			name: "empty email",
			login: func() *repository.LoginStruct {
				return &repository.LoginStruct{
					Email:    "",
					Password: password,
				}
			},
			isValid: false,
		},
		{
			name: "empty password",
			login: func() *repository.LoginStruct {
				return &repository.LoginStruct{
					Email:    email,
					Password: "",
				}
			},
			isValid: false,
		},
	}

	userId, err := s.User.Create(&models.User{
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
					assert.NoError(t, s.Login(tc.login()))
				} else {
					assert.Error(t, s.Login(tc.login()))
				}
			})
		}
	}
}
