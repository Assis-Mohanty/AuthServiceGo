package services

import (
	db "authservice/db/repository"
	"authservice/models"
)

type RolePermissionsService interface {
	AssignPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	GetRolePermissionById(id int64) (*models.RolePermission, error)
	GetAllRolePermissions() ([]*models.RolePermission, error)
}

type RolePermissionsServiceImpl struct {
	rolePermissionsRepository db.RolePermissionsRepository
}

func NewRolePermissionsService(_rolePermissionsRepository db.RolePermissionsRepository) RolePermissionsService {
	return &RolePermissionsServiceImpl{
		rolePermissionsRepository: _rolePermissionsRepository,
	}
}
func (r *RolePermissionsServiceImpl) AssignPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return r.rolePermissionsRepository.AssignPermissionToRole(roleId, permissionId)
}
func (r *RolePermissionsServiceImpl) RemovePermissionFromRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	return r.rolePermissionsRepository.RemovePermissionFromRole(roleId, permissionId)
}
func (r *RolePermissionsServiceImpl) GetRolePermissionById(id int64) (*models.RolePermission, error) {
	return r.rolePermissionsRepository.GetRolePermissionById(id)
}
func (r *RolePermissionsServiceImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	return r.rolePermissionsRepository.GetAllRolePermissions()
}
