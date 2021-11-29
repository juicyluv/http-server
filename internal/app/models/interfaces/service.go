package interfaces

import "github.com/ellywynn/http-server/internal/app/models"

type UserService interface {
	SignUp(user *models.User) (int, error)
	GetByEmail(email string) (*models.User, error)
	GetById(userId int) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	GetAll() (*[]models.User, error)
	Update(userId int) error
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
