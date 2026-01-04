package controllers

import (
	"authservice/middlewares"
	"authservice/models"
	"authservice/services"
	"authservice/utils"
	"encoding/json"
	"fmt"
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
	idStr:=r.PathValue("id")
	if idStr==""{
		http.Error(w,"missing id ",http.StatusBadRequest)
	}
	idInt,err:=strconv.ParseInt(idStr,10,64)
	if err!=nil{
		http.Error(w,"id is not a valid Integer",http.StatusBadRequest)
	}
	user,err:=uc.UserService.GetById(idInt)
	response:=map[string]any{
		"message":"Succesfully fetched users",
		"success":true,
		"error":nil,
		"data":user,
	}
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
        http.Error(w, "failed to encode response", http.StatusInternalServerError)
        return
    }
}


func (uc *UserController) Create(w http.ResponseWriter,r *http.Request){
	req,ok:=r.Context().Value(middlewares.CreateRequestkeyStruct).(*models.CreateRequestType)
	if !ok || req==nil{
		http.Error(w,"Creating user failed , invalid body",http.StatusBadRequest)
	}
	username:=req.Username
	password:=req.Password
	email:=req.Email
	if username=="" || email ==""|| password==""{
		http.Error(w,"invalid request body",http.StatusBadRequest)
		return 
	}
	uc.UserService.Create(username,email,password)
	w.Write([]byte("User creation endpoint"))	
}


func (uc *UserController) GetAllUsers(w http.ResponseWriter,r *http.Request){
	result,err:=uc.UserService.GetAllUsers()
	if err!=nil{
		fmt.Println("Fetching all users failed")
		http.Error(w,"Fetching all users failed",http.StatusBadGateway)
		response:=utils.CreateResponse("ajsndjad",false,err,nil)
		utils.WriteJson(w,http.StatusOK,response)
	}
	response:=utils.CreateResponse("Successfully fetched all users",true,nil,result)
	utils.WriteJson(w,http.StatusOK,response)

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


func (uc *UserController) GetUserByEmail(w http.ResponseWriter,r *http.Request){
	email, ok := r.Context().Value(middlewares.JwtContextKey).(string)
	if !ok || email == "" {
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Invalid or missing email in token", nil)
		return
	}

	user, err := uc.UserService.GetUserByEmail(email)
	if err != nil {
		utils.WriteJsonErrorResponse(w, http.StatusNotFound, "User not found", err)
		return
	}

	utils.WriteJsonSuccessResponse(w, http.StatusOK, "Fetched user by email successfully", user)
}

func (uc *UserController) Login(w http.ResponseWriter,r *http.Request){
	fmt.Println("asda")
	
	req,ok:=r.Context().Value(middlewares.LoginKeyStruct).(*models.LoginRequestType)
	if !ok || req==nil{
		http.Error(w,"Creating user failed , invalid body",http.StatusBadRequest)
	}
	password:=req.Password
	email:=req.Email
	fmt.Println("asdaqqq")
	data:=uc.UserService.Login(email,password)
	if data==""{
		fmt.Println("error service didnt not generate jwt")
		utils.WriteJsonErrorResponse(w, http.StatusUnauthorized, "Invalid credentials", nil)
		return
	}
	fmt.Println("asdaqqqwwwqq")
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Login Succesfull",data)
}