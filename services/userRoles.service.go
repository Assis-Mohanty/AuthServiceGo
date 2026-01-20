package services

import (
	db "authservice/db/repository"
	"authservice/models"
)

type UserRoleService interface {
	AssignRoleToUser(userId int64, roleId int64) (*models.UserRole, error)
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserRoles(userId int64) ([]*models.UserRole, error)
	GetUserPermissions(userId int64) ([]*models.Permission, error)
	HasPermission(userId int64, resource string, action string) (bool, error)
	HasRole(userId int64, roleId int64) (bool, error)
}
type UserRoleServiceImpl struct {
	userRoleRepository db.UserRoleRepository
}

func NewUserRoleService(_userRoleRepository db.UserRoleRepository) UserRoleService {
	return &UserRoleServiceImpl{
		userRoleRepository: _userRoleRepository,
	}
}

func (u *UserRoleServiceImpl) AssignRoleToUser(userId int64, roleId int64) (*models.UserRole, error) {
	return u.userRoleRepository.AssignRoleToUser(userId, roleId)
}
func (u *UserRoleServiceImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	return u.userRoleRepository.RemoveRoleFromUser(userId, roleId)
}
func (u *UserRoleServiceImpl) GetUserRoles(userId int64) ([]*models.UserRole, error) {
	return u.userRoleRepository.GetUserRoles(userId)
}
func (u *UserRoleServiceImpl) GetUserPermissions(userId int64) ([]*models.Permission, error) {
	return u.userRoleRepository.GetUserPermissions(userId)
}
func (u *UserRoleServiceImpl) HasPermission(userId int64, resource string, action string) (bool, error) {
	return u.userRoleRepository.HasPermission(userId, resource, action)
}
func (u *UserRoleServiceImpl) HasRole(userId int64, roleId int64) (bool, error) {
	return u.userRoleRepository.HasRole(userId, roleId)
}
