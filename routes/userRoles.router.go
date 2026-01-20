package routes

import (
	"authservice/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRolesRouter struct {
	userRolesController *controllers.UserRoleController
}

func NewUserRolesRouter(_userRolesController *controllers.UserRoleController) Router {
	return &UserRolesRouter{
		userRolesController: _userRolesController,
	}
}

func (urr *UserRolesRouter) Register(r *chi.Mux) {
	r.Get("/userroles/{userid}", urr.userRolesController.GetUserRoles)
	r.Get("/userpermissions/{userid}", urr.userRolesController.GetUserPermissions)
	r.Get("/userhaspermission/", urr.userRolesController.HasPermission)
	r.Get("/userhasrole/", urr.userRolesController.HasRole)
	r.Post("/userroles", urr.userRolesController.AssignRoleToUser) 
	r.Delete("/userroles/", urr.userRolesController.RemoveRoleFromUser)
}