package repository

import (
	"errors"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type AuthRepository struct {
	db             *sqlx.DB
	userRepository models.UserRepository
}

func NewAuthRepository(db *sqlx.DB, userRepo *models.UserRepository) *AuthRepository {
	return &AuthRepository{
		db:             db,
		userRepository: *userRepo,
	}
}

func (ar *AuthRepository) LogIn(input *models.AuthLoginStruct) error {
	u, err := ar.userRepository.FindByEmail(input.Email)
	if err != nil || !u.ComparePassword(input.Password) {
		return errors.New("incorrect email or password")
	}

	return nil
}

func (ar *AuthRepository) LogOut() error {
	return nil
}