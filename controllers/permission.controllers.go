package controllers

import (
	"authservice/models"
	"authservice/services"
	"authservice/utils"
	"net/http"
	"strconv"
)

type PermissionController struct {
	PermissionService services.PermissionService
}

func NewPermissionController(permissionService services.PermissionService) *PermissionController {
	return &PermissionController{
		PermissionService: permissionService,
	}
}

func (pc *PermissionController) GetPermissionByID(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "id is not a valid Integer", http.StatusBadRequest)
		return
	}
	permission, err := pc.PermissionService.GetPermissionById(idInt)
	if err != nil {
		http.Error(w, "failed to fetch permission: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, permission)
}

func (pc *PermissionController) GetAllPermissions(w http.ResponseWriter, r *http.Request) {
	permissions, err := pc.PermissionService.GetAllPermissions()
	if err != nil {
		http.Error(w, "failed to fetch permissions: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, permissions)
}
func (pc *PermissionController) CreatePermission(w http.ResponseWriter, r *http.Request) {
	var permission *models.CreatePermissionRequest
	if err:=utils.ReadJson(r,&permission); err!=nil{
		utils.WriteJson(w,http.StatusBadRequest,"Invalid request body")
		return
	}
	if permission.Name == "" || permission.Description == "" || permission.Resource == "" || permission.Action == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	createdPermission, err := pc.PermissionService.CreatePermission(permission.Name, permission.Description, permission.Resource, permission.Action)
	if err != nil {
		http.Error(w, "failed to create permission: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusCreated, createdPermission)
}

func (pc *PermissionController) DeletePermission(w http.ResponseWriter, r *http.Request) {
	id:= r.PathValue("id")
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, "id is not a valid Integer", http.StatusBadRequest)
		return
	}
	err = pc.PermissionService.DeletePermission(idInt)
	if err != nil {
		http.Error(w, "failed to delete permission: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (pc *PermissionController) UpdatePermission(w http.ResponseWriter, r *http.Request) {
	idStr:= r.PathValue("id")
	idInt, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, "id is not a valid Integer", http.StatusBadRequest)
		return
	}
	var permission *models.CreatePermissionRequest
	if err:=utils.ReadJson(r,&permission); err!=nil{
		utils.WriteJson(w,http.StatusBadRequest,"Invalid request body")
		return
	}
	if permission.Name == "" || permission.Description == "" || permission.Resource == "" || permission.Action == "" {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	updatedPermission, err := pc.PermissionService.UpdatePermission(idInt, permission.Name, permission.Description, permission.Resource, permission.Action)
	if err != nil {
		http.Error(w, "failed to update permission: "+err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteJson(w, http.StatusOK, updatedPermission)
}

