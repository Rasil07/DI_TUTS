package repository

import (
	"dependency_injection_tut/infrastructure"
	"dependency_injection_tut/model"
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