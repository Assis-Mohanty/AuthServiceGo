package routes

import (
	"authservice/controllers"

	"github.com/go-chi/chi/v5"
)

type PermissionRouter struct {
	permissionController *controllers.PermissionController
}

func NewPermissionRouter(_permissionController *controllers.PermissionController) Router {
	return &PermissionRouter{
		permissionController: _permissionController,
	}
}
func (pr *PermissionRouter) Register(r *chi.Mux) {
	r.Get("/permissions/", pr.permissionController.GetAllPermissions)
	r.Get("/permission/{id}", pr.permissionController.GetPermissionByID)
	r.Post("/permission", pr.permissionController.CreatePermission)
	r.Delete("/permission/{id}", pr.permissionController.DeletePermission)
	r.Put("/permission/{id}", pr.permissionController.UpdatePermission)
}
