package interfaces

import "github.com/ellywynn/http-server/server/internal/app/models"

type AuthRepository interface {
	LogIn(input AuthLoginStruct) (*models.User, error)
	LogOut() error
}

type UserRepository interface {
	Create(user *models.User) (int, error)
	FindByEmail(email string) (*models.User, error)
	FindByEmailWithPassword(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindById(userId int) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(userId int, user *models.UserUpdateInput) error
	Delete(userId int) error

	AddTravel(userId, travelId int) error
	GetTravels(userId int) (*[]models.Travel, error)
	RemoveTravel(userId, travelId int) error
}

type TravelRepository interface {
	Create(travel *models.Travel) (uint, error)
	FindById(travelId int) (*models.Travel, error)
	FindAll() (*[]models.Travel, error)
	Update(travelId int, travel *models.UpdateTravelInput) error
	Delete(travelId int) error
}

type PlaceRepository interface {
	Create(place *models.Place) (uint, error)
	FindById(placeId int) (*models.Place, error)
	FindAll() (*[]models.Place, error)
	Update(placeId int, place *models.UpdatePlaceInput) error
	Delete(placeId int) error
}

type UserRoleRepository interface {
	Create(role *models.UserRole) (uint, error)
	FindById(roleId int) (*models.UserRole, error)
	FindAll() (*[]models.UserRole, error)
	Update(roleId int, place *models.UpdateUserRoleInput) error
	Delete(roleId int) error
}
