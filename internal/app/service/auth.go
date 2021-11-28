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

func (as *AuthService) Login(input *models.AuthLoginStruct) error {
	return as.repository.Login(input)
}

func (as *AuthService) Logout() error {
	return as.repository.Logout()
}
