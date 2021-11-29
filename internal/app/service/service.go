package service

import (
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
	"github.com/ellywynn/http-server/internal/app/repository"
)

type Service struct {
	User interfaces.UserService
	Auth interfaces.AuthService
}

// Returns new Service instance
func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(&repo.User),
		Auth: NewAuthService(&repo.Auth),
	}
}
