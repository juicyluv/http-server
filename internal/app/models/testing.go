package models

import (
	"testing"

	"github.com/lane-c-wagner/go-tinydate"
)

// Returns default user entity
func TestUser(t *testing.T) *User {
	userRole := "User"
	return &User{
		Email:             "user@mail.com",
		Username:          "admin",
		EncryptedPassword: "qwerty",
		Role:              &userRole,
	}
}

// Returns default travel entity
func TestTravel(t *testing.T) *Travel {
	partySize := 10

	return &Travel{
		Title:        "Title",
		DurationDays: 10,
		Price:        200,
		PartySize:    &partySize,
		Complexity:   4,
		Place:        1,
		Description:  "desc",
		Date:         tinydate.Now().Format("20060102"),
	}
}
