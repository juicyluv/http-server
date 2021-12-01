package models

type UserRole struct {
	Id   uint   `json:"id"`
	Role string `json:"role"`
}

type UpdateUserRole struct {
	Role *string `json:"role"`
}
