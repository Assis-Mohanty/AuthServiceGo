package services

import (
	db "authservice/db/repository"
	"authservice/models"
)

type UserService interface {
	Create() (*models.User,error)
	GetById() (*models.User, error)
	GetAllUsers()([]*models.User,error)
	DeleteById(id int64) error
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

func (u *UserServiceImpl) Create() (*models.User,error){
	return u.userRepository.Create()
}

func (u *UserServiceImpl) GetAllUsers() ([]*models.User,error){
	return u.userRepository.GetAllUsers()
}

func (u *UserServiceImpl) DeleteById(id int64) error{
	return u.userRepository.DeleteById(id)
}
