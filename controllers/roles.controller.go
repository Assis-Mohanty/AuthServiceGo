package controllers

import (
	"authservice/models"
	"authservice/services"
	"authservice/utils"
	"net/http"
	"strconv"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: _roleService,
	}
}

func (rc *RoleController) GetRolebyId(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id ", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "id is not a valid Integer", http.StatusBadRequest)
		utils.WriteJson(w,http.StatusBadRequest,"Invalid id")
		return 
	}
	role, err := rc.roleService.GetRolebyId(idInt)
	if err != nil {
		http.Error(w, "failed to fetch role: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, role)
}
func (rc *RoleController) GetAllRoles(w http.ResponseWriter, r *http.Request) {
	roles, err := rc.roleService.GetAllRoles()
	if err != nil {
		http.Error(w, "failed to fetch roles: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, roles)
}
func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	var role *models.Role
	if err:=utils.ReadJson(r,&role); err!=nil{
		utils.WriteJson(w,http.StatusBadRequest,"Invalid request body")
	}
	name:=role.Name
	description:=role.Description
	if name == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if description == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	role, err := rc.roleService.CreateRole(name, description)
	if err != nil {
		http.Error(w, "failed to create role: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, role)
}
func (rc *RoleController) DeleteRole(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id ", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "id is not a valid Integer", http.StatusBadRequest)
		return
	}
	err = rc.roleService.DeleteRole(idInt)
	if err != nil {
		http.Error(w, "failed to delete role: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Role deleted successfully", nil)
}
func (rc *RoleController) UpdateRole(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	if idStr == "" {
		http.Error(w, "missing id ", http.StatusBadRequest)
		return
	}
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "id is not a valid Integer", http.StatusBadRequest)
		return
	}
	var role *models.Role
	if err:=utils.ReadJson(r,&role); err!=nil{
		utils.WriteJson(w,http.StatusBadRequest,"Invalid request body")
	}
	name:=role.Name
	description:=role.Description
	if name == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	if description == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	updatedRole, err := rc.roleService.UpdateRole(idInt, name, description)
	if err != nil {
		http.Error(w, "failed to update role: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, updatedRole)
}