package app

import (
	dbConfig "authservice/config/db"
	"authservice/controllers"
	repo "authservice/db/repository"
	"authservice/routes"
	"authservice/services"
	"fmt"
	"net/http"
	"time"

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
	urr:=repo.NewUserRoleRepository(db)
	us:=services.NewUserService(ur,urr)
	uc:=controllers.NewUserController(us)
	uRouter:=routes.NewUserRouter(uc)
	rr:=repo.NewRoleRepository(db)
	rs:=services.NewRoleService(rr)
	rc:=controllers.NewRoleController(rs)
	rRouter:=routes.NewRoleRouter(rc)
	pr:=repo.NewPermissionRepository(db)
	ps:=services.NewPermissionService(pr)
	pc:=controllers.NewPermissionController(ps)
	pRouter:=routes.NewPermissionRouter(pc)
	urs:=services.NewUserRoleService(urr)
	urc:=controllers.NewUserRoleController(urs)
	urRouter:=routes.NewUserRolesRouter(urc)

	server:=&http.Server{
		Addr: app.Config.Address,
		Handler: routes.SetUpRouter(uRouter,rRouter,pRouter,urRouter),
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10 *time.Second,
	}

	return server.ListenAndServe()
}


