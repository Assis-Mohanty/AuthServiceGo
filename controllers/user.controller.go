package controllers

import (
	"authservice/models"
	"authservice/services"
	"authservice/utils"
	"encoding/json"
	"fmt"

	// "log"
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

// func (uc *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {
// 	log.Println("aksnmdkam")
//     idStr := r.PathValue("id")
//     if idStr == "" {
//         http.Error(w, "missing id", http.StatusBadRequest)
//         return
//     }

//     idInt, err := strconv.ParseInt(idStr, 10, 64)
//     if err != nil {
//         http.Error(w, "id must be a valid integer", http.StatusBadRequest)
//         return
//     }

//     user, err := uc.UserService.GetById(idInt)
//     if err != nil {
//         http.Error(w, "user not found", http.StatusNotFound)
//         return
//     }

//     w.Header().Set("Content-Type", "application/json")
//     w.WriteHeader(http.StatusOK)

//     if err := json.NewEncoder(w).Encode(user); err != nil {
//         // at this point headers are already sent, just log
//         log.Println("encode error:", err)
//         return
//     }
// }


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
	uc.UserService.GetUserByEmail(email)
	w.Write([]byte("Fetching all users"))
}

func (uc *UserController) Login(w http.ResponseWriter,r *http.Request){
	var req models.LoginRequestType
	// if err:=json.NewDecoder(r.Body).Decode(&req);err !=nil{
	// 	http.Error(w,"invalid request body",http.StatusBadRequest)
	// }
	jsonErr:=utils.ReadJson(r ,&req); 
	if jsonErr !=nil{
		w.Write([]byte("Failed reading Json"))
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Failed reading Json",jsonErr)
		return
	}

	if validatorErr:=utils.Validator.Struct(req);validatorErr!=nil{
		w.Write([]byte("Invalid input data"))
		utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Invalid input data",jsonErr)
		return
	}
	password:=req.Password
	email:=req.Email
	data:=uc.UserService.Login(email,password)
	utils.WriteJsonSuccessResponse(w,http.StatusOK,"Login Succesfull",data)
	
}