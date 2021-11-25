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
