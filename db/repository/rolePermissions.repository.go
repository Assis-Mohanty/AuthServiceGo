package db

import (
	"authservice/models"
	"database/sql"
)

type RolePermissionsRepository interface {
	AssignPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRole(roleId int64, permissionId int64) (*models.RolePermission, error)
	GetRolePermissionById(id int64) (*models.RolePermission, error)
	GetAllRolePermissions() ([]*models.RolePermission, error)
}

type RolePermissionsRepositoryImpl struct {
	db *sql.DB
}
func NewRolePermissionsRepository(db *sql.DB) RolePermissionsRepository {
	return &RolePermissionsRepositoryImpl{
		db: db,
	}
}

func (r *RolePermissionsRepositoryImpl) GetRolePermissionById(id int64) (*models.RolePermission, error) {
	query:= "SELECT id,role_id,permission_id,created_at,updated_at FROM role_permissions WHERE id=?"
	row:= r.db.QueryRow(query,id)
	rolePermission:=&models.RolePermission{}
	if err:=row.Scan(&rolePermission.Id,&rolePermission.RoleId,&rolePermission.PermissionId,&rolePermission.CreatedAt,&rolePermission.UpdatedAt);err!=nil{
		return nil,err
	}
	return rolePermission,nil
}
func (r *RolePermissionsRepositoryImpl) AssignPermissionToRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	query:= "INSERT INTO role_permissions (role_id,permission_id) VALUES (?,?)"
	result,err:= r.db.Exec(query,roleId,permissionId)
	if err!=nil{
		return nil,err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		return nil,err
	}
	return r.GetRolePermissionById(id)
}
func (r *RolePermissionsRepositoryImpl) RemovePermissionFromRole(roleId int64, permissionId int64) (*models.RolePermission, error) {
	query:= "DELETE FROM role_permissions WHERE role_id=? AND permission_id=?"
	_,err:= r.db.Exec(query,roleId,permissionId)
	if err!=nil{
		return nil,err
	}
	return &models.RolePermission{},nil
}
func (r *RolePermissionsRepositoryImpl) GetAllRolePermissions() ([]*models.RolePermission, error) {
	query:= "SELECT id,role_id,permission_id,created_at,updated_at FROM role_permissions"
	rows,err:= r.db.Query(query)
	if err!=nil{
		return nil,err
	}
	defer rows.Close()

	var rolePermissions []*models.RolePermission
	for rows.Next() {
		rolePermission:=&models.RolePermission{}
		if err:=rows.Scan(&rolePermission.Id,&rolePermission.RoleId,&rolePermission.PermissionId,&rolePermission.CreatedAt,&rolePermission.UpdatedAt);err!=nil{
			return nil,err
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}
	return rolePermissions, nil
}
