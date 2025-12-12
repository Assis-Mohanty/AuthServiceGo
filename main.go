package main
import(
	app "authservice/app"
)

func main(){
	cnf:=app.NewConfig(":3000")
	app:=app.NewApplication(cnf)
	app.Run()
}