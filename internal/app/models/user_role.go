package models

type UserRole struct {
	Id   uint   `json:"id"`
	Role string `json:"role"`
}

type UpdateUserRoleInput struct {
	Role *string `json:"role"`
}
