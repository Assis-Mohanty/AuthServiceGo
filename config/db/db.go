package config

import (
	"authservice/config"
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func SetUpDb()(*sql.DB,error){
	cfg:=mysql.NewConfig()
	cfg.User=config.GetString("DB_USER","root")
	cfg.Passwd=config.GetString("DB_PASSWORD","qqqq")
	cfg.Net=config.GetString("DB_NETWORK","tcp")
	cfg.Addr=config.GetString("DB_ADDRESS","127.0.0.1:3306")
	cfg.DBName=config.GetString("DB_NAME","usersgo")
	cfg.ParseTime=true
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