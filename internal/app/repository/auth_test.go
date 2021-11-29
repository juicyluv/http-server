package repository_test

import (
	"testing"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
	"github.com/ellywynn/http-server/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepository_LogIn(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("users")

	testUser := models.TestUser(t)
	password := testUser.EncryptedPassword

	userId, err := r.User.Create(testUser)

	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)

	u, err := r.Auth.LogIn(&interfaces.AuthLoginStruct{
		Email:    testUser.Email,
		Password: password,
	})

	assert.NoError(t, err)
	assert.NotNil(t, u)
}
