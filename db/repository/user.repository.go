package db

import (
	"authservice/models"
	"database/sql"
	"fmt"
)

// import "database/sql"

type UserRepository interface {
	Create() (*models.User,error)
	GetById() (*models.User,error)
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository{
	return &UserRepositoryImpl{
		db:_db,
	}
}

func (u *UserRepositoryImpl) GetById() (*models.User ,error) {
	fmt.Println("Fetching User in UserRepository")
	query:="SELECT id,username,email,created_at,updated_at FROM users WHERE id=?"
	row:=u.db.QueryRow(query,1)
	user:=&models.User{}
	err:=row.Scan(&user.Id,&user.Username,&user.Email,&user.CreatedAt,&user.UpdatedAt)

	if err!= nil{
		if err==sql.ErrNoRows{
			fmt.Println("No user found with the given Id")
			return nil,err
		}else{
			fmt.Println("Error scanning user:",err)
			return nil,err
		}
	}
	fmt.Println("User fetched succesfully:",user)
	return user,nil
}

func (u *UserRepositoryImpl) Create()(*models.User,error){
	query:="INSERT INTO users (username,email,password) VALUES (?,?,?)"
	user:=&models.User{}
	result, err := u.db.Exec(
		query,
		"kakdsnmkamsd","aisdalsqsda","asdlmalmsdl")
	if err!=nil{
		return nil,err
	}
	id,err:=result.LastInsertId()
	if err!=nil{
		return nil,err
	}
	fmt.Println("User Created with id:",id)
	user.Id=id
	return user,nil
}