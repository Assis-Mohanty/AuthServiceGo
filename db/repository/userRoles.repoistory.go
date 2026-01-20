package db

import (
	"authservice/models"
	"database/sql"
)
type UserRoleRepository interface {
	AssignRoleToUser(userId int64, roleId int64) (*models.UserRole, error)
	RemoveRoleFromUser(userId int64, roleId int64) error
	GetUserRoles(userId int64) ([]*models.UserRole, error)
	GetUserPermissions(userId int64) ([]*models.Permission, error)
	HasPermission(userId int64, resource string, action string) (bool, error)
	HasRole(userId int64, roleId int64) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
}
type UserRoleRepositoryImpl struct {
	db *sql.DB
}
func NewUserRoleRepository(db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: db,
	}
}

func (ur *UserRoleRepositoryImpl) GetUserRoleById(id int64) (*models.UserRole, error) {
	query := "SELECT id, user_id, role_id, created_at, updated_at FROM user_roles WHERE id = ?"
	row := ur.db.QueryRow(query, id)
	userRole := &models.UserRole{}
	if err := row.Scan(&userRole.Id, &userRole.UserId, &userRole.RoleId, &userRole.CreatedAt, &userRole.UpdatedAt); err != nil {
		return nil, err
	}
	return userRole, nil
}

func (ur *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) (*models.UserRole, error) {
	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	result, err := ur.db.Exec(query, userId, roleId)
	if err != nil {
		return nil, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	return ur.GetUserRoleById(id)
}

func (ur *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	_, err := ur.db.Exec(query, userId, roleId)
	return err
}
func (ur *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.UserRole, error) {
	query := "SELECT id, user_id, role_id, created_at, updated_at FROM user_roles WHERE user_id = ?"
	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []*models.UserRole
	for rows.Next() {
		userRole := &models.UserRole{}
		if err := rows.Scan(&userRole.Id, &userRole.UserId, &userRole.RoleId, &userRole.CreatedAt, &userRole.UpdatedAt); err != nil {
			return nil, err
		}
		userRoles = append(userRoles, userRole)
	}
	return userRoles, nil
}

func (ur *UserRoleRepositoryImpl) GetUserPermissions(userId int64) ([]*models.Permission, error) {
	query := `
	SELECT p.id, p.name, p.description, p.resource, p.action, p.created_at, p.updated_at
	FROM permissions p
	JOIN role_permissions rp ON p.id = rp.permission_id
	JOIN user_roles ur ON rp.role_id = ur.role_id
	WHERE ur.user_id = ?`

	rows, err := ur.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var permissions []*models.Permission
	for rows.Next() {
		permission := &models.Permission{}
		if err := rows.Scan(&permission.Id, &permission.Name, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (ur *UserRoleRepositoryImpl) HasPermission(userId int64, resource string, action string) (bool, error) {
	query := `
	SELECT COUNT(*)
	FROM permissions p
	JOIN role_permissions rp ON p.id = rp.permission_id
	JOIN user_roles ur ON rp.role_id = ur.role_id
	WHERE ur.user_id = ? AND p.resource = ? AND p.action = ?`
	row := ur.db.QueryRow(query, userId, resource, action)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}
func (ur *UserRoleRepositoryImpl) HasRole(userId int64, roleId int64) (bool, error) {
	query := `
	SELECT COUNT(*)
	FROM user_roles
	WHERE user_id = ? AND role_id = ?`
	row := ur.db.QueryRow(query, userId, roleId)
	var count int
	if err := row.Scan(&count); err != nil {
		return false, err
	}
	return count > 0, nil
}

func (ur *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {
	query := `
	SELECT COUNT(DISTINCT r.name)
	FROM roles r
	JOIN user_roles ur ON r.id = ur.role_id
	WHERE ur.user_id = ? AND r.name IN  (?)
	Group BY ur.user_id`
	row := ur.db.QueryRow(query, len(roleNames), userId, roleNames)
	var hasAllRoles bool
	if err := row.Scan(&hasAllRoles); err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return hasAllRoles, nil
}
