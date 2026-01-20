package controllers

import (
	"authservice/models"
	"authservice/services"
	"authservice/utils"
	"net/http"
	"strconv"
)

type UserRoleController struct {
	userRoleService services.UserRoleService
}

func NewUserRoleController(_userRoleService services.UserRoleService) *UserRoleController {
	return &UserRoleController{
		userRoleService: _userRoleService,
	}
}

func (urc *UserRoleController) AssignRoleToUser(w http.ResponseWriter, r *http.Request) {
	var req *models.UserRoleRequest
	err:= utils.ReadJson(r, &req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userId := req.UserId
	roleId := req.RoleId
	if userId == 0 || roleId == 0 {
		http.Error(w, "userId and roleId must be provided", http.StatusBadRequest)
		return
	}
	userRole, err := urc.userRoleService.AssignRoleToUser(userId, roleId)
	if err != nil {
		http.Error(w, "failed to assign role to user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, userRole)
}

func (urc *UserRoleController) RemoveRoleFromUser(w http.ResponseWriter, r *http.Request) {
	var req *models.UserRoleRequest
	err:= utils.ReadJson(r, &req)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userId := req.UserId
	roleId := req.RoleId
	if userId == 0 || roleId == 0 {
		http.Error(w, "userId and roleId must be provided", http.StatusBadRequest)
		return
	}
	err = urc.userRoleService.RemoveRoleFromUser(userId, roleId)
	if err != nil {
		http.Error(w, "failed to remove role from user: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (urc *UserRoleController) GetUserRoles(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userid")
	if userIdStr == "" {
		http.Error(w, "userid query parameter is required", http.StatusBadRequest)
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid userid", http.StatusBadRequest)
		return
	}
	userRoles, err := urc.userRoleService.GetUserRoles(userId)
	if err != nil {
		http.Error(w, "failed to get user roles: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, userRoles)
}

func (urc *UserRoleController) GetUserPermissions(w http.ResponseWriter, r *http.Request) {
	userIdStr := r.PathValue("userid")
	if userIdStr == "" {
		http.Error(w, "userid query parameter is required", http.StatusBadRequest)
		return
	}
	userId, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid userid", http.StatusBadRequest)
		return
	}
	permissions, err := urc.userRoleService.GetUserPermissions(userId)
	if err != nil {
		http.Error(w, "failed to get user permissions: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, permissions)
}
func (urc *UserRoleController) HasPermission(w http.ResponseWriter, r *http.Request) {
	var userRole *models.HasPermissionRequest
	err:= utils.ReadJson(r, &userRole)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userId := userRole.UserId
	resource := userRole.Resource
	action := userRole.Action
	if userId == 0 || resource == "" || action == "" {
		http.Error(w, "userId, resource and action must be provided", http.StatusBadRequest)
		return
	}
	hasPermission, err := urc.userRoleService.HasPermission(userId, resource, action)
	if err != nil {
		http.Error(w, "failed to check permission: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]bool{"has_permission": hasPermission})
}
func (urc *UserRoleController) HasRole(w http.ResponseWriter, r *http.Request) {
	var userRole *models.HasRoleRequest
	err:= utils.ReadJson(r, &userRole)
	if err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	userId := userRole.UserId
	roleId := userRole.RoleId
	if userId == 0 || roleId == 0 {
		http.Error(w, "userId and roleId must be provided", http.StatusBadRequest)
		return
	}
	hasRole, err := urc.userRoleService.HasRole(userId, roleId)
	if err != nil {
		http.Error(w, "failed to check role: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, map[string]bool{"has_role": hasRole})
}


