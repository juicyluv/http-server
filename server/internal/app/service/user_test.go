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

	userId, err := s.User.SignUp(models.TestUser(t))
	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)
}

func TestUserService_FindByEmail(t *testing.T) {
	s, teardown := service.NewTestService(t, dbURL)
	defer teardown("users")

	email := "user@example.org"
	userId, err := s.User.SignUp(&models.User{
		Email:             email,
		Username:          "user",
		EncryptedPassword: "qwerty",
	})

	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)

	user, err := s.User.GetByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)

	noUser, err := s.User.GetByEmail("another@gmail.com")
	assert.Error(t, err)
	assert.Nil(t, noUser)
}

func TestUserService_GetByUsername(t *testing.T) {
	s, teardown := service.NewTestService(t, dbURL)
	defer teardown("users")

	testUser := models.TestUser(t)

	userId, err := s.User.SignUp(testUser)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)

	user, err := s.User.GetByUsername(testUser.Username)
	assert.NotNil(t, user)
	assert.NoError(t, err)

	noUser, err := s.User.GetByUsername("invalid nickname")
	assert.Nil(t, noUser)
	assert.Error(t, err)
}
