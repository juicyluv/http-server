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

func (r *UserRepository) Create(u *models.User) (int, error) {
	if err := u.HashPassword(); err != nil {
		return 0, err
	}

	var userId int
	query := "INSERT INTO users (email, username, encrypted_password) VALUES ($1, $2, $3) RETURNING id"
	err := r.db.QueryRow(query, u.Email, u.Username, u.EncryptedPassword).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	query := "SELECT id, email, username FROM users WHERE email = $1"
	err := r.db.QueryRow(query, email).Scan(&u.Id, &u.Username, &u.Email)
	if err != nil {
		return nil, err
	}

	return u, nil
}
