package repository

import (
	"errors"
	"fmt"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/ellywynn/http-server/server/internal/app/models/interfaces"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db             *sqlx.DB
	userRepository interfaces.UserRepository
}

func NewAuthRepository(db *sqlx.DB, userRepo *interfaces.UserRepository) *AuthRepository {
	return &AuthRepository{
		db:             db,
		userRepository: *userRepo,
	}
}

func (ar *AuthRepository) LogIn(input interfaces.AuthLoginStruct) (*models.User, error) {
	u, err := ar.userRepository.FindByEmailWithPassword(input.Email)
	if err != nil || !u.ComparePassword(input.Password) {
		fmt.Println(err.Error())
		return nil, errors.New("incorrect email or password")
	}

	return u, nil
}

func (ar *AuthRepository) LogOut() error {
	return nil
}
