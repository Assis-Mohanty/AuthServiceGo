package app

import (
	"fmt"
	"net/http"
	"time"
)

type Config struct{
	Address string
}

type Application struct{
	Config Config
}
func NewConfig(address string) Config{
	return Config{
		Address: address,
	}
}

func NewApplication(config Config)*Application{
	return &Application{
		Config: config,
	}
}

func (app *Application) Run() error{
	server:=&http.Server{
		Addr: app.Config.Address,
		Handler: nil,
		ReadTimeout: 10 *time.Second,
		WriteTimeout: 10 *time.Second,
	}
	fmt.Println("Starting server on",app.Config.Address)
	return server.ListenAndServe()
}