package services

import (
	db "authservice/db/repository"
	"authservice/models"
	"os"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(username string,email string,password string) (*models.User,error)
	GetById() (*models.User, error)
	GetAllUsers()([]*models.User,error)
	DeleteById(id int64) error
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

func (u *UserServiceImpl) GetById() (*models.User, error){
	return u.userRepository.GetById()
}

func (u *UserServiceImpl) Create(username string,email string,password string) (*models.User,error){
	hashedPassword,_:=bcrypt.GenerateFromPassword([]byte(password),bcrypt.DefaultCost)
	return u.userRepository.Create(username,email,string(hashedPassword))

}



func (u *UserServiceImpl) GetAllUsers() ([]*models.User,error){
	return u.userRepository.GetAllUsers()
}

func (u *UserServiceImpl) DeleteById(id int64) error{
	return u.userRepository.DeleteById(id)
}

func GenerateJWT() string{
	jwt:=jwt.New(jwt.SigningMethodES256)
	jwtString,err:=jwt.SignedString(os.Getenv("JWT_KEY"))
	if err!=nil{
		return err.Error()
	}
	return jwtString 
}
