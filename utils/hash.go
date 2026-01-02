package utils

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func GenerateHashPassword(password string) (string, error) {
	hashedPassword, err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err!=nil{
		return "",err
	}
	return string(hashedPassword),err
}

func CheckPasswordHash(hashedPassword string ,password string )bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPassword),[]byte(password))
	if err!=nil{
		fmt.Println("Password do not match the hashedPassword")
		return false
	}
	return true 
}