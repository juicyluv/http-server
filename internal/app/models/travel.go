package models

import "time"

type Travel struct {
	Id           uint      `json:"id"`
	Title        string    `json:"title" binding:"required"`
	DurationDays int       `json:"duration_days" db:"duration_days" binding:"required"`
	Price        int       `json:"price" binding:"required"`
	PartySize    int       `json:"party_size" db:"party_size"`
	Complexity   int       `json:"complexity" binding:"required"`
	Place        int       `json:"place" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Date         time.Time `json:"date" binding:"required"`
}
