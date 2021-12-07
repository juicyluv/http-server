package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
)

type AuthService struct {
	repository interfaces.AuthRepository
}

func NewAuthService(repo *interfaces.AuthRepository) interfaces.AuthService {
	return &AuthService{
		repository: *repo,
	}
}

func (as *AuthService) LogIn(input interfaces.AuthLoginStruct) (*models.User, error) {
	return as.repository.LogIn(input)
}

func (as *AuthService) LogOut() error {
	return as.repository.LogOut()
}
