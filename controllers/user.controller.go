package controllers

import (
	"authservice/models"
	"authservice/services"
	"authservice/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type CreateRequestType struct{
	Username string `json:"username"`
	Email string `json:"email" validate:"required,email"`
	Password string `json:"password"`

}

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
	var req CreateRequestType
	if err:=json.NewDecoder(r.Body).Decode(&req);err !=nil{
		http.Error(w,"invalid request body",http.StatusBadRequest)
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
	email:=r.PathValue("email")
	user,err:=uc.UserService.GetUserByEmail(email)
	if err!=nil{
		w.Write([]byte("Fetching by email failed"))
	}
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Fetched user by email succesfully",user)
	w.Write([]byte("Fetching all users"))
}

func (uc *UserController) Login(w http.ResponseWriter,r *http.Request){
	fmt.Println("asda")
	var req models.LoginRequestType
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