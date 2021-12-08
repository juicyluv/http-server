package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	db *sqlx.DB
}

// Create new User Repository instance
func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

// Create user and return new user ID
func (r *UserRepository) Create(u *models.User) (int, error) {
	if err := u.HashPassword(); err != nil {
		return 0, err
	}

	// If user role presented, find its id
	var role int
	if u.Role != nil {
		err := r.db.Get(&role, "SELECT id FROM user_roles WHERE role = $1", *u.Role)
		if err != nil {
			// If user role is invalid, return an error
			if err == sql.ErrNoRows {
				return 0, errors.New("there is no role " + *u.Role)
			}
			return 0, err
		}
	}

	valuesQuery := "(DEFAULT, $1, $2, $3, $4)"
	args := make([]interface{}, 0)
	args = append(args, u.Email, u.Username, u.EncryptedPassword)

	// If role is not presented, use default value
	// Otherwise, add role to arguments
	if role == 0 {
		valuesQuery = "(DEFAULT, $1, $2, $3)"
	} else {
		args = append(args, role)
	}

	var userId int
	query := fmt.Sprintf("INSERT INTO users VALUES %s RETURNING id", valuesQuery)
	err := r.db.QueryRow(query, args...).Scan(&userId)
	if err != nil {
		return 0, err
	}

	return userId, nil
}

// Find user by email and return User struct
func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	u := models.User{}
	query := `
	SELECT u. id, u. email, u.username, ur.role AS role 
	FROM users u
	INNER JOIN user_roles ur
	ON u.role = ur.id
	WHERE u.email=$1`
	if err := r.db.Get(&u, query, email); err != nil {
		return nil, err
	}

	return &u, nil
}

// Find user by email and return User struct with user password
// Use it carefully
func (r *UserRepository) FindByEmailWithPassword(email string) (*models.User, error) {
	u := models.User{}
	query := `
	SELECT u.id, u.email, u.username, u.encrypted_password, ur.role AS role 
	FROM users u
	INNER JOIN user_roles ur
	ON u.role = ur.id
	WHERE u.email=$1`
	if err := r.db.Get(&u, query, email); err != nil {
		return nil, err
	}

	return &u, nil
}

// Find user by username and return user instance
func (r *UserRepository) FindByUsername(username string) (*models.User, error) {
	u := models.User{}
	query := `
	SELECT u. id, u. email, u.username, ur.role AS role 
	FROM users u
	INNER JOIN user_roles ur
	ON u.role = ur.id
	WHERE u.username=$1`
	if err := r.db.Get(&u, query, username); err != nil {
		return nil, err
	}

	return &u, nil
}

// Find user by Id and return user instance
func (r *UserRepository) FindById(userId int) (*models.User, error) {
	u := models.User{}
	query := `
	SELECT u. id, u. email, u.username, ur.role AS role 
	FROM users u
	INNER JOIN user_roles ur
	ON u.role = ur.id
	WHERE u.id=$1`
	if err := r.db.Get(&u, query, userId); err != nil {
		return nil, err
	}

	return &u, nil
}

// Return all users
func (r *UserRepository) GetAll() (*[]models.User, error) {
	u := []models.User{}
	query := `
	SELECT u. id, u. email, u.username, ur.role AS role 
	FROM users u
	INNER JOIN user_roles ur
	ON u.role = ur.id`
	if err := r.db.Select(&u, query); err != nil {
		return nil, err
	}

	return &u, nil
}

// Update the user
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

	if user.Role != nil {
		values = append(values, fmt.Sprintf("role=$%d", argId))
		args = append(args, *user.Role)
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
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	query := "DELETE FROM users_travels WHERE user_id = $1"
	_, err = r.db.Exec(query, userId)
	if err != nil {
		return err
	}

	query = "DELETE FROM users WHERE id = $1"
	_, err = tx.Exec(query, userId)
	if err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

// Add user travel
func (r *UserRepository) AddTravel(userId, travelId int) error {
	query := "INSERT INTO users_travels VALUES ($1, $2)"
	_, err := r.db.Exec(query, userId, travelId)
	return err
}

// Get user travels by user id
func (r *UserRepository) GetTravels(userId int) (*[]models.Travel, error) {
	var travelsIds []string
	var travels []models.Travel

	idsQuery := "SELECT travel_id FROM users_travels WHERE user_id=$1"
	if err := r.db.Select(&travelsIds, idsQuery, userId); err != nil {
		return nil, err
	}

	if len(travelsIds) == 0 {
		return &travels, nil
	}

	travelsQuery := strings.Join(travelsIds, ", ")
	query := fmt.Sprintf("SELECT * FROM travels WHERE id IN (%s)", travelsQuery)
	if err := r.db.Select(&travels, query); err != nil {
		return nil, err
	}

	return &travels, nil
}

// Remove user travel
func (r *UserRepository) RemoveTravel(userId, travelId int) error {
	query := "DELETE FROM users_travels WHERE user_id = $1 AND travel_id = $2"
	_, err := r.db.Exec(query, userId, travelId)
	return err
}

// hash user password
func hashPassword(password string) (string, error) {
	// Hash password
	b, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(b), nil
}
