package main

import (
	app "authservice/app"
	"authservice/config"
)

func main(){
	config.Load()
	cnf:=app.NewConfig(":3000")
	app:=app.NewApplication(cnf)
	app.Run()
}