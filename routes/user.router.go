package routes

import (
	"authservice/controllers"
	"authservice/middlewares"

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
	r.Get("/profile/{id}",ur.userController.GetUserById)
	r.Post("/create",ur.userController.Create)
	r.Get("/getallprofiles",ur.userController.GetAllUsers)
	r.Delete("/profile/{id}",ur.userController.DeleteById)
	r.Get("/getbyemail/{email}",ur.userController.GetUserByEmail)
	r.With(middlewares.UserLoginMiddleware).Post("/verify",ur.userController.Login)
}
