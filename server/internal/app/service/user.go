package service

import (
	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/ellywynn/http-server/server/internal/app/models/interfaces"
)

type UserService struct {
	repository interfaces.UserRepository
}

func NewUserService(repo *interfaces.UserRepository) interfaces.UserService {
	return &UserService{
		repository: *repo,
	}
}

func (us *UserService) SignUp(user *models.User) (int, error) {
	return us.repository.Create(user)
}

func (us *UserService) GetByEmail(email string) (*models.User, error) {
	return us.repository.FindByEmail(email)
}

func (us *UserService) GetByUsername(username string) (*models.User, error) {
	return us.repository.FindByUsername(username)
}

func (us *UserService) GetById(userId int) (*models.User, error) {
	return us.repository.FindById(userId)
}

func (us *UserService) GetAll() (*[]models.User, error) {
	return us.repository.GetAll()
}

func (us *UserService) Update(userId int, user *models.UserUpdateInput) error {
	return us.repository.Update(userId, user)
}

func (us *UserService) Delete(userId int) error {
	return us.repository.Delete(userId)
}

func (us *UserService) AddTravel(userId, travelId int) error {
	return us.repository.AddTravel(userId, travelId)
}

func (us *UserService) GetTravels(userId int) (*[]models.Travel, error) {
	return us.repository.GetTravels(userId)
}

func (us *UserService) RemoveTravel(userId, travelId int) error {
	return us.repository.RemoveTravel(userId, travelId)
}
