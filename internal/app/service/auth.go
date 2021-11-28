package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
)

type AuthService struct {
	repository models.AuthRepository
}

func NewAuthService(repo *models.AuthRepository) models.AuthService {
	return &AuthService{
		repository: *repo,
	}
}

func (as *AuthService) LogIn(input *models.AuthLoginStruct) (*models.User, error) {
	return as.repository.LogIn(input)
}

func (as *AuthService) LogOut() error {
	return as.repository.LogOut()
}
