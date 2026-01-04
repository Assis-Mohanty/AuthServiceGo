package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type jwtContextKey struct{}

var JwtContextKey =jwtContextKey{}

func JwtVerifyMiddleware(next http.Handler)http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader:=r.Header.Get("Authorization")
		if authHeader==""{
			http.Error(w,"Authorization header is required",http.StatusUnauthorized)
			return
		}
		bearerToken:="Bearer "
		if !strings.HasPrefix(authHeader,bearerToken){
			http.Error(w,"Bearer token is required",http.StatusUnauthorized)
			return 
		}
		token:=strings.TrimPrefix(authHeader,bearerToken)
		if token==""{
			http.Error(w,"JWT token is missing",http.StatusUnauthorized)
			return 
		}
		claims:=jwt.MapClaims{}
		_,err:= jwt.ParseWithClaims(token,claims,func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_KEY")),nil
		})
		if err!=nil{
			http.Error(w,"Invalid token:"+err.Error(),http.StatusUnauthorized)
			return
		}
		email,ok:=claims["email"].(string)
		if!ok {
			http.Error(w,"Invalid token claims",http.StatusUnauthorized)
			return 
		}
		cxt:=context.WithValue(r.Context(),JwtContextKey,email)
		next.ServeHTTP(w,r.WithContext(cxt))


	})
}