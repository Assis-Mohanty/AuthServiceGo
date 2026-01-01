package controllers

import (
	"authservice/services"
	"net/http"
	"strconv"
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

func (uc *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	uc.UserService.GetAllUsers()
	w.Write([]byte("Fetching all users"))
}

func (uc *UserController) DeleteById(w http.ResponseWriter,r *http.Request){
	idStr:=r.PathValue("id")
	if idStr==""{
		http.Error(w,"missing id ",http.StatusBadRequest)
	}
	idInt,err:=strconv.ParseInt(idStr,10,64)
	if err!=nil{
		http.Error(w,"id is not a valid Integer",http.StatusBadRequest)
	}
	uc.UserService.DeleteById(idInt)
	w.Write([]byte("Deleting a user"))
}


