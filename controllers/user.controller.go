package controllers

import (
	"authservice/services"
	"net/http"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController(_userService services.UserService) *UserController{
	return &UserController{
		UserService: _userService,
	}
}

func (uc *UserController) GetUserById(w http.ResponseWriter,r *http.Request){
	uc.UserService.GetById()
	w.Write([]byte("User registration endpoint")) 
}

func (uc *UserController) Create(w http.ResponseWriter,r *http.Request){
	uc.UserService.Create()
	w.Write([]byte("User creation endpoint"))
}