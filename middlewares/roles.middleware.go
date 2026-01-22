package middlewares

import (
	"authservice/config"
	db "authservice/db/repository"
	"authservice/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func RequireAllRoles(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIdStr:=chi.URLParam(r,"id")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}
		dbConn, dbErr := config.SetUpDb()
		if dbErr != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		urr := db.NewUserRoleRepository(dbConn)
		hasAllRoles, hasAllRolesErr := urr.HasAllRoles(userId, roles)
		if hasAllRolesErr != nil {
			http.Error(w, "Error checking user roles:"+hasAllRolesErr.Error(), http.StatusInternalServerError)
			return
		}
		if !hasAllRoles {
			http.Error(w, "Forbidden: insufficient roles", http.StatusForbidden)
			return
		}
		fmt.Println("User has all required roles:", roles)
		next.ServeHTTP(w, r)
		})
	}
}



func RequireAnyRole(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIdStr:=chi.URLParam(r,"id")
		userId, err := strconv.ParseInt(userIdStr, 10, 64)
		if err != nil {
			http.Error(w, "Invalid user ID", http.StatusUnauthorized)
			return
		}
		dbConn, dbErr := config.SetUpDb()
		if dbErr != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		urr := db.NewUserRoleRepository(dbConn)
		hasAnyRole, hasAnyRoleErr := urr.HasAnyRole(userId, roles)
		if hasAnyRoleErr != nil {
			http.Error(w, "Error checking user roles:"+hasAnyRoleErr.Error(), http.StatusInternalServerError)
			return
		}
		if !hasAnyRole {
			http.Error(w, "Forbidden: insufficient roles", http.StatusForbidden)
			return
		}
		fmt.Println("User has any required roles:", roles)
		next.ServeHTTP(w, r)
		})
	}
}


func RequireAnyRoleJwtAuth(roles ...string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// userIdStr:=chi.URLParam(r,"id")
		// jwtContextKey:=middlewares.JwtContextKey
		email,ok:=r.Context().Value(utils.JwtContextKey).(string)
		if !ok || email==""{
			http.Error(w, "Invalid or missing email in token", http.StatusUnauthorized)
			return
		}
		
		dbConn, dbErr := config.SetUpDb()
		if dbErr != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		ur:=db.NewUserRepository(dbConn)
		user, userErr:=ur.GetUserByEmail(email)
		if userErr!=nil{
			http.Error(w, "Error fetching user by email:"+userErr.Error(), http.StatusInternalServerError)
			return
		}
		userId:=user.Id
		// userId, err := strconv.ParseInt(userIdstr, 10, 64)
		// if err != nil {
		// 	http.Error(w, "Invalid user ID", http.StatusUnauthorized)
		// 	return
		// }
		urr := db.NewUserRoleRepository(dbConn)
		hasAnyRole, hasAnyRoleErr := urr.HasAnyRole(userId, roles)
		if hasAnyRoleErr != nil {
			http.Error(w, "Error checking user roles:"+hasAnyRoleErr.Error(), http.StatusInternalServerError)
			return
		}
		if !hasAnyRole {
			http.Error(w, "Forbidden: insufficient roles", http.StatusForbidden)
			return
		}
		fmt.Println("User has any required roles:", roles)
		next.ServeHTTP(w, r)
		})
	}
}
