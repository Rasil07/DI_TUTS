package repository

import (
	"dependency_injection_tut/infrastructure"
	"dependency_injection_tut/model"
	"fmt"
)


type UserRepository struct{
	DB *infrastructure.Database
}

func NewUserRepository(db *infrastructure.Database) *UserRepository{
	return &UserRepository{DB:db}
}

func(ur *UserRepository) CreateUser(user *model.User) (*model.User,error) {
return user, ur.DB.Create(user).Error
}

func(ur *UserRepository) CheckUserPresent(user *model.User) (*model.User,error){
	fmt.Println(user.Email)
	return user, ur.DB.First(&user,"email = ?", user.Email ).Error
}