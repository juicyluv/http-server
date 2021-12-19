package service

import (
	"github.com/ellywynn/http-server/server/internal/app/models/interfaces"
	"github.com/ellywynn/http-server/server/internal/app/repository"
)

type Service struct {
	User     interfaces.UserService
	Auth     interfaces.AuthService
	Travel   interfaces.TravelService
	Place    interfaces.PlaceService
	UserRole interfaces.UserRoleService

	Cld interfaces.CloudinaryService
}

// Returns new Service instance
func NewService(repo *repository.Repository) *Service {
	return &Service{
		User:     NewUserService(&repo.User),
		Auth:     NewAuthService(&repo.Auth),
		Travel:   NewTravelService(&repo.Travel),
		Place:    NewPlaceService(&repo.Place),
		UserRole: NewUserRoleService(&repo.UserRole),

		Cld: NewCloudinaryService(),
	}
}
