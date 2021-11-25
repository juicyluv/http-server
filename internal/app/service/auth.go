package service

import "github.com/ellywynn/http-server/internal/app/repository"

type AuthService struct {
	repository repository.Auth
}

type Auth interface {
	Login(input *repository.LoginStruct) error
	Logout() error
}

func NewAuthService(repo repository.Auth) *AuthService {
	return &AuthService{
		repository: repo,
	}
}

func (as *AuthService) Login(input *repository.LoginStruct) error {
	return as.repository.Login(input)
}

func (as *AuthService) Logout() error {
	return as.repository.Logout()
}
