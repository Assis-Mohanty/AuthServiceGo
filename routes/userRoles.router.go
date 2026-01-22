package routes

import (
	"authservice/controllers"
	"authservice/middlewares"

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
	r.With(middlewares.JwtVerifyMiddleware,middlewares.RequireAnyRoleJwtAuth("admin")).Post("/userroles", urr.userRolesController.AssignRoleToUser)
	r.Delete("/userroles", urr.userRolesController.RemoveRoleFromUser)
	r.Get("/userroles/{userid}", urr.userRolesController.GetUserRoles)
	r.Get("/userpermissions/{userid}", urr.userRolesController.GetUserPermissions)
	r.Post("/userhaspermission", urr.userRolesController.HasPermission)
	r.Post("/userhasrole", urr.userRolesController.HasRole)
}
