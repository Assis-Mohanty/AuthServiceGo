package models

import (
	"database/sql"
	"time"
)

type Role struct {
	Id          int64
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   sql.NullTime
}

type Permission struct {
	Id          int64
	Name        string
	Description string
	Resource    string
	Action      string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserRole struct {
	Id        int64
	UserId    int64
	RoleId    int64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type RolePermission struct {
	Id           int64
	RoleId       int64
	PermissionId int64
	CreatedAt    time.Time
	UpdatedAt    time.Time
}


type CreatePermissionRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Resource    string `json:"resource"`
	Action      string `json:"action"`
}
type UserRoleRequest struct {
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
}
type HasPermissionRequest struct {
	UserId   int64  `json:"user_id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
type HasRoleRequest struct {
	UserId int64 `json:"user_id"`
	RoleId int64 `json:"role_id"`
}
type AssignPermissionToRoleRequest struct {
	RoleId int64 `json:"role_id"`
	PermissionId int64 `json:"permission_id"`
}