package routes

import (
	"authservice/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	userController *controllers.UserController
}

func NewUserRouter(_userController *controllers.UserController) Router{
	return &UserRouter{
		userController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Mux){
	r.Get("/profile",ur.userController.GetUserById)
	r.Post("/create",ur.userController.Create)
}
