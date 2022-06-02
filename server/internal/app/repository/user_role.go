package repository

import (
	"fmt"
	"strings"

	"github.com/ellywynn/http-server/server/internal/app/models"
	"github.com/ellywynn/http-server/server/internal/app/models/interfaces"
	"github.com/jmoiron/sqlx"
)

type UserRoleRepository struct {
	db *sqlx.DB
}

func NewUserRoleRepository(db *sqlx.DB) interfaces.UserRoleRepository {
	return &UserRoleRepository{
		db: db,
	}
}

func (ur *UserRoleRepository) Create(role *models.UserRole) (uint, error) {
	var roleId uint
	query := "INSERT INTO user_roles (role) VALUES ($1) RETURNING id"
	if err := ur.db.QueryRow(query, role.Role).Scan(&roleId); err != nil {
		return 0, err
	}

	return roleId, nil
}

func (ur *UserRoleRepository) FindAll() (*[]models.UserRole, error) {
	var roles []models.UserRole
	query := "SELECT * FROM user_roles"
	if err := ur.db.Select(&roles, query); err != nil {
		return nil, err
	}

	return &roles, nil
}

func (ur *UserRoleRepository) FindById(roleId int) (*models.UserRole, error) {
	var role models.UserRole
	query := "SELECT * FROM user_roles WHERE id = $1"
	if err := ur.db.Get(&role, query, roleId); err != nil {
		return nil, err
	}

	return &role, nil
}

func (ur *UserRoleRepository) Update(roleId int, role *models.UpdateUserRoleInput) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if role.Role != nil {
		values = append(values, fmt.Sprintf("role=%d", argId))
		args = append(args, role.Role)
		argId++
	}

	valuesQuery := strings.Join(values, ", ")
	args = append(args, roleId)

	query := fmt.Sprintf("UPDATE user_roles SET %s WHERE id=$%d", valuesQuery, argId)
	_, err := ur.db.Exec(query, args...)

	return err
}

func (ur *UserRoleRepository) Delete(roleId int) error {
	query := "DELETE FROM user_roles WHERE id = $1"
	_, err := ur.db.Exec(query, roleId)

	return err
}
