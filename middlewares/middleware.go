package middlewares

import (
	"authservice/models"
	"authservice/utils"
	"context"
	"net/http"
)

type loginkeyStruct struct{}
var LoginKeyStruct =loginkeyStruct{}

type createRequestkeyStruct struct{}
var  CreateRequestkeyStruct=createRequestkeyStruct{}

func UserLoginMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter,r *http.Request){
		req:=&models.LoginRequestType{}
		if err:=utils.ReadJson(r,req);err!=nil{
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Failed Reading Json",err)
			return
		}
		if validatorErr:=utils.Validator.Struct(req);validatorErr!=nil{
			w.Write([]byte("Validation failed"))
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Validation failed",validatorErr)
			return
		}
		cxt:=context.WithValue(r.Context(),LoginKeyStruct,req)
		next.ServeHTTP(w,r.WithContext(cxt))
	})
}

func UserCreateMiddleware(next http.Handler) http.Handler{
	return http.HandlerFunc(func (w http.ResponseWriter,r *http.Request){
		req:=&models.CreateRequestType{}
		if err:=utils.ReadJson(r,req);err!=nil{
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Failed Reading Json",err)
			return
		}
		if validatorErr:=utils.Validator.Struct(req);validatorErr!=nil{
			w.Write([]byte("Validation failed"))
			utils.WriteJsonErrorResponse(w,http.StatusBadRequest,"Validation failed",validatorErr)
			return
		}
		cxt:=context.WithValue(r.Context(),CreateRequestkeyStruct,req)
		next.ServeHTTP(w,r.WithContext(cxt))
	})
}


