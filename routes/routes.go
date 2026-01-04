package routes

import (
	"authservice/controllers"
	"authservice/middlewares"
	"authservice/utils"

	// "authservice/middlewares"

	"github.com/go-chi/chi/v5"
)

type Router interface{
	Register(r *chi.Mux)
}

func SetUpRouter(UserRouter Router) *chi.Mux {
	router:=chi.NewRouter()
	router.Use(middlewares.RateLimiterMiddleware)
	router.HandleFunc("/fakestoreapiservice/*",utils.ReverseProxy("https://fakestoreapi.com","/fakestoreapiservice"))
	router.Get("/ping", controllers.PingHandler)
	UserRouter.Register(router)
	return router
}