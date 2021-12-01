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
}

type AuthService interface {
	LogIn(input *AuthLoginStruct) (*models.User, error)
	LogOut() error
}

type AuthLoginStruct struct {
	Email    string
	Password string
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
