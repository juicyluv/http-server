package models

type UserService interface {
	SignUp(user *User) (int, error)
	GetByEmail(email string) (*User, error)
	GetById(userId int) (*User, error)
	GetByUsername(username string) (*User, error)
	GetAll() (*[]User, error)
	Update(userId int) error
	Delete(userId int) error
}

type UserRepository interface {
	Create(user *User) (int, error)
	FindByEmail(email string) (*User, error)
	FindByUsername(username string) (*User, error)
	FindById(userId int) (*User, error)
	GetAll() (*[]User, error)
	Update(userId int) error
	Delete(userId int) error
}

type AuthRepository interface {
	LogIn(input *AuthLoginStruct) error
	LogOut() error
}

type AuthService interface {
	Login(input *AuthLoginStruct) error
	Logout() error
}

type AuthLoginStruct struct {
	Email    string
	Password string
}
