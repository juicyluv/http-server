package repository

import (
	"fmt"
	"strings"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
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
	u := models.User{}
	query := "SELECT id, email, username FROM users WHERE email=$1"
	if err := r.db.Get(&u, query, email); err != nil {
		return nil, err
	}

	return &u, nil
}

// Finds user by email and returns User struct with user password
// Use it carefully
func (r *UserRepository) FindByEmailWithPassword(email string) (*models.User, error) {
	u := models.User{}
	query := "SELECT id, email, username, encrypted_password FROM users WHERE email=$1"
	if err := r.db.Get(&u, query, email); err != nil {
		return nil, err
	}

	return &u, nil
}

// Finds user by username and returns user instance
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	u := models.User{}
	query := "SELECT id, email, username FROM users WHERE username=$1"
	if err := r.db.Get(&u, query, username); err != nil {
		return nil, err
	}

	return &u, nil
}

// Finds user by Id and returns user instance
func (r *UserRepository) FindById(userId int) (*models.User, error) {
	u := models.User{}
	query := "SELECT id, email, username FROM users WHERE id=$1"
	if err := r.db.Get(&u, query, userId); err != nil {
		return nil, err
	}

	return &u, nil
}

// Returns all users
func (r *UserRepository) GetAll() (*[]models.User, error) {
	u := []models.User{}
	query := "SELECT id, email, username FROM users"
	if err := r.db.Select(&u, query); err != nil {
		return nil, err
	}

	return &u, nil
}

// Updates the user
func (r *UserRepository) Update(userId int, user *models.UserUpdateInput) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if user.Email != nil {
		values = append(values, fmt.Sprintf("email=$%d", argId))
		args = append(args, *user.Email)
		argId++
	}

	if user.Username != nil {
		values = append(values, fmt.Sprintf("username=$%d", argId))
		args = append(args, *user.Username)
		argId++
	}

	if user.Password != nil {
		// Hash user password
		hashedPassword, err := hashPassword(*user.Password)
		if err != nil {
			return err
		}
		values = append(values, fmt.Sprintf("encrypted_password=$%d", argId))
		args = append(args, hashedPassword)
		argId++
	}

	valuesQuery := strings.Join(values, ", ")
	query := fmt.Sprintf("UPDATE users SET %s WHERE id = $%d", valuesQuery, argId)
	args = append(args, userId)

	_, err := r.db.Exec(query, args...)
	return err
}

// Deletes the user
func (r *UserRepository) Delete(userId int) error {
	query := "DELETE FROM users WHERE id = $1"
	_, err := r.db.Exec(query, userId)
	return err
}

func hashPassword(password string) (string, error) {
	// Hash password
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
