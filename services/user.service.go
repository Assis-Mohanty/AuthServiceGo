package services

import (
	db "authservice/db/repository"
	"authservice/models"
)

type UserService interface {
	Create() (*models.User,error)
	GetById() (*models.User, error)
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