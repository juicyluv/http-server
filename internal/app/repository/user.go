package repository

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

// Creates new User Repository instance
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Creates user in database and returns new user ID
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

// Finds user by email and returns User struct
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := &models.User{}
	query := "SELECT id, email, username, encrypted_password FROM users WHERE email=$1"
	err := r.db.QueryRow(query, email).Scan(&u.Id, &u.Username, &u.Email, &u.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Funds user by username and return user instance
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	u := &models.User{}
	query := "SELECT id, email, username, encrypted_password FROM users WHERE username=$1"
	err := r.db.QueryRow(query, username).Scan(&u.Id, &u.Username, &u.Email, &u.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Funds user by Id and return user instance
func (r *UserRepository) FindById(userId int) (*models.User, error) {
	u := &models.User{}
	query := "SELECT id, email, username, ecnrypted_password FROM users WHERE id=$1"
	err := r.db.QueryRow(query, userId).Scan(&u.Id, &u.Username, &u.Email, &u.EncryptedPassword)
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Returns all users
func (r *UserRepository) GetAll() (*[]models.User, error) {
	u := &[]models.User{}
	query := "SELECT id, email, username FROM users"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		rows.Scan(&user.Id, &user.Email, &user.Username)
		*u = append(*u, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return u, nil
}

// Updates the user
func (r *UserRepository) Update(userId int) error {
	return nil
}

// Deletes the user
func (r *UserRepository) Delete(userId int) error {
	return nil
}
