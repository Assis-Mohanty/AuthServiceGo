package services

import (
	db "authservice/db/repository"
	"authservice/models"
)

type PermissionService interface {
	GetAllPermissions() ([]*models.Permission, error)
	GetPermissionById(id int64) (*models.Permission, error)
	CreatePermission(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermission(id int64) error
	UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

type PermissionServiceImpl struct {
	permissionRepository db.PermissionRepository
}

func NewPermissionService(_permissionRepository db.PermissionRepository) PermissionService {
	return &PermissionServiceImpl{
		permissionRepository: _permissionRepository,
	}
}
func (p *PermissionServiceImpl) GetAllPermissions() ([]*models.Permission, error) {
	return p.permissionRepository.GetAllPermissions()
}
func (p *PermissionServiceImpl) GetPermissionById(id int64) (*models.Permission, error) {
	return p.permissionRepository.GetPermissionById(id)
}
func (p *PermissionServiceImpl) CreatePermission(name string, description string, resource string, action string) (*models.Permission, error) {
	return p.permissionRepository.CreatePermission(name, description, resource, action)
}
func (p *PermissionServiceImpl) DeletePermission(id int64) error {
	return p.permissionRepository.DeletePermission(id)
}
func (p *PermissionServiceImpl) UpdatePermission(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	return p.permissionRepository.UpdatePermission(id, name, description, resource, action)
}

