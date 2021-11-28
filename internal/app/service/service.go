package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/repository"
)

type Service struct {
	User models.UserService
	Auth models.AuthService
}

// Returns new Service instance
func NewService(repo *repository.Repository) *Service {
	return &Service{
		User: NewUserService(&repo.User),
		Auth: NewAuthService(&repo.Auth),
	}
}
