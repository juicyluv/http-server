package models

type Place struct {
	Id   uint   `json:"id"`
	Name string `json:"name" binding:"required"`
}
