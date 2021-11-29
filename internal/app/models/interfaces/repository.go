package interfaces

import "github.com/ellywynn/http-server/internal/app/models"

type AuthRepository interface {
	LogIn(input *AuthLoginStruct) (*models.User, error)
	LogOut() error
}

type UserRepository interface {
	Create(user *models.User) (int, error)
	FindByEmail(email string) (*models.User, error)
	FindByUsername(username string) (*models.User, error)
	FindById(userId int) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(userId int, user *models.UserUpdateInput) error
	Delete(userId int) error
}

type TravelRepository interface {
	Create(travel *models.Travel) (int, error)
	FindById(travelId int) (*models.Travel, error)
	FindAll() (*[]models.Travel, error)
	Update(travel *models.Travel) error
	Delete(travelId int) error
}
