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
	GetAllUsers()([]*models.User,error)
	DeleteById(id int64) error
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


func (u *UserRepositoryImpl) GetAllUsers()([] *models.User,error){
	query:="SELECT id,username,email FROM users"
	rows,err:=u.db.Query(query)
	if err!=nil{
		fmt.Println("Fetching all users")
		return nil,err
	}
	defer rows.Close()
	var users[] *models.User
	for rows.Next(){
		user:=&models.User{}
		err:=rows.Scan(&user.Id,&user.Username,&user.Email)
		if err!=nil{
			return nil,err
		}
		users = append(users, user)
	}
	if err:=rows.Err();err!=nil{
		return nil,err
	}
	for _, u := range users {
	fmt.Printf("%+v\n", *u)
	}
	return users,err
}

func (u *UserRepositoryImpl) DeleteById(id int64) error{
	query:=`delete from users where id=?`
	result,err:=u.db.Exec(query,id)
	if err!=nil{
		return err
	}
	userId,err:=result.LastInsertId()
	if err!=nil{
		return err
	}
	fmt.Println("USer deleted with id:",userId)
	return nil
}
