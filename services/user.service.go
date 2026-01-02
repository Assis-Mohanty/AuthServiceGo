package services

import (
	db "authservice/db/repository"
	"authservice/models"
	"fmt"
	"os"

	"authservice/utils"

	"github.com/golang-jwt/jwt/v5"
)

type UserService interface {
	Create(username string,email string,password string) (*models.User,error)
	GetById(id int64) (*models.ResponseUserDTO, error)
	GetAllUsers()([]*models.User,error)
	DeleteById(id int64) error
	GetUserByEmail(email string )(*models.User,error)
	VerifyEmailAndPassword(email string ,password string) string
	// GenerateJWT() string
}

type UserServiceImpl struct {
	userRepository db.UserRepository
}

func NewUserService(_userRepository db.UserRepository) UserService{
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

func (u *UserServiceImpl) GetById(id int64) (*models.ResponseUserDTO, error){
	return u.userRepository.GetById(id)
}

func (u *UserServiceImpl) Create(username string,email string,password string) (*models.User,error){
	 hashedPassword,_:=utils.GenerateHashPassword(password)
	return u.userRepository.Create(username,email,string(hashedPassword))

}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User,error){
	return u.userRepository.GetAllUsers()
}

func (u *UserServiceImpl) DeleteById(id int64) error{
	return u.userRepository.DeleteById(id)
}

func GenerateJWT() string{
	jwt:=jwt.New(jwt.SigningMethodHS256)
	string:=os.Getenv("JWT_KEY")
	jwtString,err:=jwt.SignedString([]byte(string) )
	if err!=nil{
		return err.Error()
	}
	return jwtString
}

func (u *UserServiceImpl) GetUserByEmail(email string )(*models.User,error){
	return u.userRepository.GetUserByEmail(email)
}

func (u *UserServiceImpl) VerifyEmailAndPassword(email string ,password string) string{
	user,err:=u.userRepository.GetUserByEmail(email)
	if err==nil{
		fmt.Println("user fetched successfully ",user.Id,user.Username,user.Email,user.Password)
	}
	result:=utils.CheckPasswordHash(user.Password,password)
	if result{
		fmt.Println("Password is Correct")
		jwt:=GenerateJWT()
		fmt.Println(jwt)
		return string(jwt)
	}
	return ""

}
