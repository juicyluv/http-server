package interfaces

import "github.com/ellywynn/http-server/internal/app/models"

type UserService interface {
	SignUp(user *models.User) (int, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(userId int, user *models.UserUpdateInput) error
	Delete(userId int) error

	AddTravel(userId, travelId int) error
	GetTravels(userId int) (*[]models.Travel, error)
	RemoveTravel(userId, travelId int) error
}

type AuthService interface {
	LogIn(input AuthLoginStruct) (*models.User, error)
	LogOut() error
}

type AuthLoginStruct struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TravelService interface {
	Create(travel *models.Travel) (uint, error)
	GetById(travelId int) (*models.Travel, error)
	GetAll() (*[]models.Travel, error)
	Update(travelId int, travel *models.UpdateTravelInput) error
	Delete(travelId int) error
}

type PlaceService interface {
	Create(place *models.Place) (uint, error)
	GetById(placeId int) (*models.Place, error)
	GetAll() (*[]models.Place, error)
	Update(placeId int, place *models.UpdatePlaceInput) error
	Delete(placeId int) error
}

type UserRoleService interface {
	Create(role *models.UserRole) (uint, error)
	GetById(roleId int) (*models.UserRole, error)
	GetAll() (*[]models.UserRole, error)
	Update(roleId int, place *models.UpdateUserRoleInput) error
	Delete(roleId int) error
}

type CloudinaryService interface {
	UploadImage(image, name, folder string) (string, error)
}
