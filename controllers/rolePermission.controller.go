package controllers

import (
	"authservice/models"
	"authservice/services"
	"authservice/utils"
	"net/http"
	"strconv"
)

type RolePermissionController struct {
	rolePermissionService services.RolePermissionsService
}

func NewRolePermissionController(rolePermissionService services.RolePermissionsService) *RolePermissionController {
	return &RolePermissionController{
		rolePermissionService: rolePermissionService,
	}
}
func (rpc *RolePermissionController) AssignPermissionToRole(w http.ResponseWriter, r *http.Request) {
	var req *models.AssignPermissionToRoleRequest
	err:= utils.ReadJson(r, &req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	roleId := req.RoleId
	permissionId := req.PermissionId
	if roleId == 0 || permissionId == 0 {
		http.Error(w, "roleId and permissionId must be provided", http.StatusBadRequest)
		return
	}
	rolePermission, err := rpc.rolePermissionService.AssignPermissionToRole(roleId, permissionId)
	if err != nil {
		http.Error(w, "failed to assign permission to role: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK,rolePermission)
}
func (rpc *RolePermissionController) RemovePermissionFromRole(w http.ResponseWriter, r *http.Request) {
	var req *models.AssignPermissionToRoleRequest
	err:= utils.ReadJson(r, &req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	roleId := req.RoleId
	permissionId := req.PermissionId
	if roleId == 0 || permissionId == 0 {
		http.Error(w, "roleId and permissionId must be provided", http.StatusBadRequest)
		return
	}
	rolePermission, err := rpc.rolePermissionService.RemovePermissionFromRole(roleId, permissionId)
	if err != nil {
		http.Error(w, "failed to remove permission from role: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK,rolePermission)
}
func (rpc *RolePermissionController) GetRolePermissionById(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "id is not a valid Integer", http.StatusBadRequest)
		return
	}
	rolePermission, err := rpc.rolePermissionService.GetRolePermissionById(idInt)
	if err != nil {
		http.Error(w, "failed to fetch role permission: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, rolePermission)
}
func (rpc *RolePermissionController) GetAllRolePermissions(w http.ResponseWriter, r *http.Request) {
	rolePermissions, err := rpc.rolePermissionService.GetAllRolePermissions()
	if err != nil {
		http.Error(w, "failed to fetch role permissions: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, rolePermissions)
}

