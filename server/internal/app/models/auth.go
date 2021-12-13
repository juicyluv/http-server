package models

type LoginResponse struct {
	UserId   int    `json:"id"`
	Email    string `json:"email"`
	Username string `json:"username"`
	Role     string `json:"role"`
}
