package repository

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(u *models.User) (*models.User, error) {
	return nil, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	return nil, nil
}
