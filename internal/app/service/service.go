package service

import "github.com/ellywynn/http-server/internal/app/repository"

type Service struct {
	User
	Auth
	// Repo *repository.Repository
}

// Returns new Service instance
func NewService(repo *repository.Repository) *Service {
	return &Service{
		// Repo: repo,
		User: NewUserService(repo),
		Auth: NewAuthService(repo),
	}
}
