package service

import "github.com/ellywynn/http-server/internal/app/repository"

type Service struct {
	User
	Auth
}

// Returns new Service instance
func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(repo),
		Auth: NewAuthService(repo),
	}
}
