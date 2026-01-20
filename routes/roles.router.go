package routes

import (
	"authservice/controllers"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	roleController *controllers.RoleController
}

func NewRoleRouter(_roleController *controllers.RoleController) Router {
	return &RoleRouter{
		roleController: _roleController,
	}
}
func (rr *RoleRouter) Register(r *chi.Mux) {
	r.Get("/roles", rr.roleController.GetAllRoles)
	r.Get("/role/{id}", rr.roleController.GetRolebyId)
	r.Post("/role", rr.roleController.CreateRole)
	r.Delete("/role/{id}", rr.roleController.DeleteRole)
	r.Put("/role/{id}", rr.roleController.UpdateRole)
}