package service_test

import (
	"testing"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/service"
	"github.com/stretchr/testify/assert"
)

func TestUserService_Create(t *testing.T) {
	s, teardown := service.NewTestService(t, dbURL)
	defer teardown("users")

	userId, err := s.User.Create(models.TestUser(t))
	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)
}

func TestUserService_FindByEmail(t *testing.T) {
	s, teardown := service.NewTestService(t, dbURL)
	defer teardown("users")

	email := "user@example.org"
	userId, err := s.User.Create(&models.User{
		Email:             email,
		Username:          "user",
		EncryptedPassword: "qwerty",
	})

	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)

	user, err := s.User.FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	noUser, err := s.User.FindByEmail("another@gmail.com")
	assert.Error(t, err)
	assert.Nil(t, noUser)
}
