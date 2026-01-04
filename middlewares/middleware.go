package middlewares

import (
	"authservice/models"
	"authservice/utils"
	"net/http"
)
var req *models.LoginRequestType
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
		next.ServeHTTP(w,r)
	})
}