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
	FindById(userId int) (*models.User, error)
	GetAllUsers() (*[]models.User, error)
	UpdateUser(userId int) error
	DeleteUser(userId int) error
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

func (us *UserService) FindById(userId int) (*models.User, error) {
	return us.repository.FindById(userId)
}

func (us *UserService) GetAllUsers() (*[]models.User, error) {
	return us.repository.GetAllUsers()
}

func (us *UserService) UpdateUser(userId int) error {
	return us.repository.UpdateUser(userId)
}

func (us *UserService) DeleteUser(userId int) error {
	return us.repository.DeleteUser(userId)
}
