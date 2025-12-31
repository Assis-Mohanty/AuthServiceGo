package config

import (
	// "database/sql"
	env "authservice/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetUpDb()(*sql.DB,error){
	cfg:=mysql.NewConfig()
	cfg.User=env.GetString("DB_USER","root")
	cfg.Passwd=env.GetString("DB_PASSWORD","qqqq")
	cfg.Net=env.GetString("DB_NETWORK","tcp")
	cfg.Addr=env.GetString("DB_ADDRESS","127.0.0.1:3306")
	cfg.DBName=env.GetString("DB_NAME","usersgo")
	var err error
	fmt.Println(cfg.FormatDSN())
	db,err = sql.Open("mysql",cfg.FormatDSN())
	if err != nil{
		log.Fatal(err)
		return nil,err
	}
	pingErr:=db.Ping()
	if pingErr !=nil{
		log.Fatal(pingErr)
		return nil, pingErr
	}
	return db,nil
}