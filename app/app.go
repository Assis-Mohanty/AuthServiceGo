package app

import (
	"authservice/controllers"
	repo "authservice/db/repository"
	"authservice/routes"
	"authservice/services"
	"fmt"
	"net/http"
	"time"
	dbConfig "authservice/config/db"
)

type Config struct{
	Address string
}

type Application struct{
	Config Config
	Storage repo.Storage
}
func NewConfig(address string) Config{
	return Config{
		Address: address,
	}
}

func NewApplication(config Config)*Application{
	return &Application{
		Config: config,
		Storage: *repo.NewStorage(),
	}
}

func (app *Application) Run() error{
	db,err:=dbConfig.SetUpDb()
	if err!=nil{
		fmt.Println("Error setting up database:",err)
		return err
	}
	fmt.Println("Starting server on",app.Config.Address)
	ur:=repo.NewUserRepository(db)
	fmt.Println("qqqq")

	us:=services.NewUserService(ur)
	uc:=controllers.NewUserController(us)
	uRouter:=routes.NewUserRouter(uc)
	server:=&http.Server{
		Addr: app.Config.Address,
		Handler: routes.SetUpRouter(uRouter),
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10 *time.Second,
	}
	return server.ListenAndServe()
}


