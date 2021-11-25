package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/repository"
)

type UserService struct {
	repository *repository.Repository
}

type UserServiceInterface interface {
	Create(user *models.User) (int, error)
	FindByEmail(email string) (*models.User, error)
}

func NewUserService(repo *repository.Repository) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (us *UserService) Create(user *models.User) (int, error) {
	userId, err := us.repository.UserRepository.Create(user)
	if err != nil {
		return 0, err
	}

	return userId, nil
}
