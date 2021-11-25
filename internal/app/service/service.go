package service

import "github.com/ellywynn/http-server/internal/app/repository"

type Service struct {
	Repository *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Repository: repo,
	}
}
