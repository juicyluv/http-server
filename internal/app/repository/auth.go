package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db             *sqlx.DB
	userRepository User
}

type Auth interface {
	Login(input *LoginStruct) error
	Logout() error
}

type LoginStruct struct {
	Email    string
	Password string
}

func NewAuthRepository(db *sqlx.DB, userRepo *User) *AuthRepository {
	return &AuthRepository{
		db:             db,
		userRepository: *userRepo,
	}
}

func (ar *AuthRepository) Login(input *LoginStruct) error {
	u, err := ar.userRepository.FindByEmail(input.Email)
	if err != nil || !u.ComparePassword(input.Password) {
		return errors.New("incorrect email or password")
	}

	return nil
}

func (ar *AuthRepository) Logout() error {
	return nil
}
