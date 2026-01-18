package services

import (
	db "authservice/db/repository"
	"authservice/models"
)

type RoleService interface {
	GetAllRoles() ([]*models.Role, error)
	GetRolebyId(id int64) (*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRole(id int64) error
	UpdateRole(id int64, name string, description string) (*models.Role, error)
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
}
func NewRoleService(_roleRepository db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: _roleRepository,
	}
}

func (r *RoleServiceImpl) GetAllRoles() ([]*models.Role, error) {
	return r.roleRepository.GetAllRoles()
}
func (r *RoleServiceImpl) GetRolebyId(id int64) (*models.Role, error) {
	return r.roleRepository.GetRolebyId(id)
}
func (r *RoleServiceImpl) CreateRole(name string, description string) (*models.Role, error) {
	return r.roleRepository.CreateRole(name, description)
}
func (r *RoleServiceImpl) DeleteRole(id int64) error {
	return r.roleRepository.DeleteRole(id)
}
func (r *RoleServiceImpl) UpdateRole(id int64, name string, description string) (*models.Role, error) {
	return r.roleRepository.UpdateRole(id, name, description)
}