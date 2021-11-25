package models_test

import (
	"testing"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/stretchr/testify/assert"
)

func TestUser_HashPassword(t *testing.T) {
	u := models.TestUser(t)
	assert.NoError(t, u.HashPassword())
	assert.NotEmpty(t, u.EncryptedPassword)
}
