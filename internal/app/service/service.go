package service

import "github.com/ellywynn/http-server/internal/app/repository"

type Service struct {
	UserService *UserService
}

// Returns new Service instance
func NewService(repo *repository.Repository) *Service {
	return &Service{
		UserService: NewUserService(repo),
	}
}
