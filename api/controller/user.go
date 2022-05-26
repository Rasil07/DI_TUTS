package controller

import (
	"dependency_injection_tut/model"
	"dependency_injection_tut/service"
	"dependency_injection_tut/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{
	userServie *service.UserService
}

type CreateUserCredentials struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
}
type UserReturnType struct{
	ID uint `json:"id"`
	NAME string `json:"name"`
	EMAIL string	`json:"email"`
}

func NewUserController(us *service.UserService) *UserController{
	return &UserController{userServie: us}
}

func (uc *UserController) ReturnUser(u *model.User) *UserReturnType {
	rtrnUser := &UserReturnType{
		ID:u.ID,
		EMAIL:u.Email,
		NAME:u.Name,
	}
return rtrnUser
}

func(uc *UserController) Create(c *gin.Context){
	var reqUser CreateUserCredentials
	if err:= c.Bind(&reqUser);err!=nil{
		log.Println("Binding context to user faile")
		return 
	}
	hashedPassword,err:= utils.HashPassword(reqUser.Password)

	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	}

	user := &model.User{
		Name: reqUser.Name,
		Email: reqUser.Email,
		Password: []byte(hashedPassword),
	}
	user,err = uc.userServie.Create(user)
	if err!=nil{
		log.Println("Binding context to user failed")
	}
	returnUser:= uc.ReturnUser(user)


	c.JSON(200,gin.H{
		"data":returnUser,
	})

}