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

func SetUpRouter(routes ...Router) *chi.Mux {
	router:=chi.NewRouter()
	router.Use(middlewares.RateLimiterMiddleware)
	router.HandleFunc("/fakestoreapiservice/*",utils.ReverseProxy("https://fakestoreapi.com","/fakestoreapiservice"))
	router.HandleFunc("/hotelservice/*",utils.ReverseProxy("http://localhost:3000","/hotelservice"))
	router.Get("/ping", controllers.PingHandler)
	for _, route := range routes {
		route.Register(router)
	}
	return router
}

