package models

import "time"

type User struct {
	Id        int64
	Username  string
	Email     string
	Password  string  `json:"-"`
	CreatedAt time.Time 
	UpdatedAt time.Time
}


type ResponseUserDTO struct {
	Id        int64
	Username  string
	Email     string
	CreatedAt time.Time 
	UpdatedAt time.Time
}