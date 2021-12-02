package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Travel struct {
	Id           uint        `json:"id"`
	Title        string      `json:"title" binding:"required"`
	DurationDays int         `json:"duration_days" db:"duration_days" binding:"required"`
	Price        int         `json:"price" binding:"required"`
	PartySize    *int        `json:"party_size" db:"party_size"`
	Complexity   int         `json:"complexity" binding:"required"`
	Place        interface{} `json:"place" binding:"required"`
	Description  string      `json:"description" binding:"required"`
	Date         time.Time   `json:"date" binding:"required"`
}

type UpdateTravelInput struct {
	Title        *string    `json:"title"`
	DurationDays *int       `json:"duration_days" db:"duration_days"`
	Price        *int       `json:"price"`
	PartySize    *int       `json:"party_size" db:"party_size"`
	Complexity   *int       `json:"complexity"`
	Place        *int       `json:"place"`
	Description  *string    `json:"description"`
	Date         *time.Time `json:"date"`
}

// Validates creating user struct
func (t *Travel) Validate() error {
	return validation.ValidateStruct(
		t,
		validation.Field(&t.Title, validation.Required),
		validation.Field(&t.DurationDays, validation.Required),
		validation.Field(&t.Price, validation.Required),
		validation.Field(&t.PartySize, is.Int),
		validation.Field(&t.Complexity, validation.Required),
		validation.Field(&t.Description, is.ASCII, validation.Required),
		validation.Field(&t.Date, validation.Required),
		validation.Field(&t.Place, validation.Required),
	)
}
