package service

import (
	"github.com/ellywynn/http-server/internal/app/models"
	"github.com/ellywynn/http-server/internal/app/models/interfaces"
)

type UserRoleService struct {
	repository interfaces.UserRoleRepository
}

func NewUserRoleService(repo *interfaces.UserRoleRepository) interfaces.UserRoleService {
	return &UserRoleService{
		repository: *repo,
	}
}

func (us *UserRoleService) Create(role *models.UserRole) (uint, error) {
	return us.repository.Create(role)
}

func (us *UserRoleService) GetAll() (*[]models.UserRole, error) {
	return us.repository.FindAll()
}

func (us *UserRoleService) GetById(roleId int) (*models.UserRole, error) {
	return us.repository.FindById(roleId)
}

func (us *UserRoleService) Update(roleId int, role *models.UpdateUserRoleInput) error {
	return us.repository.Update(roleId, role)
}

func (us *UserRoleService) Delete(roleId int) error {
	return us.repository.Delete(roleId)
}
