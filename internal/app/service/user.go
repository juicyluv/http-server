package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/repository"
)

type UserService struct {
	repository repository.User
}

type User interface {
	Create(user *models.User) (int, error)
	FindByEmail(email string) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
}

func NewUserService(repo repository.User) *UserService {
	return &UserService{
		repository: repo,
	}
}

func (us *UserService) Create(user *models.User) (int, error) {
	return us.repository.Create(user)
}

func (us *UserService) FindByEmail(email string) (*models.User, error) {
	return us.repository.FindByEmail(email)
}

func (us *UserService) GetAllUsers() (*[]models.User, error) {
	return us.repository.GetAllUsers()
}
