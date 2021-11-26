package repository_test

import (
	"testing"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/repository"
	"github.com/stretchr/testify/assert"
)

func TestUserRepository_Create(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("users")

	userId, err := r.User.Create(models.TestUser(t))

	assert.NoError(t, err)
	assert.NotEqual(t, 0, userId)
}

func TestUserRepository_FindByEmail(t *testing.T) {
	r, teardown := repository.NewTestRepository(t, dbURL)
	defer teardown("users")

	email := "user@example.com"
	_, err := r.User.FindByEmail(email)

	assert.Error(t, err)

	u := models.TestUser(t)
	u.Email = email
	_, err = r.User.Create(u)

	if err != nil {
		t.Fatalf("Error while creating the user: %s", err.Error())
	}

	user, err := r.User.FindByEmail(email)
	assert.NoError(t, err)
	assert.NotNil(t, user)
}
