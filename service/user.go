package service

import (
	"dependency_injection_tut/model"
	"dependency_injection_tut/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(urp *repository.UserRepository) *UserService{
	return &UserService{userRepository: urp}
}

func(us *UserService) Create(user *model.User) (*model.User,error){
	return us.userRepository.CreateUser(user)
}

func(us *UserService) CheckUserExists(user *model.User) (*model.User,error){
	return us.userRepository.CheckUserPresent(user)
}