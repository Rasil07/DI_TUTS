package service

import (
	"dependency_injection_tut/infrastructure"

	"github.com/gin-gonic/gin"
)

type GoogleService struct{
	infrastructure.Database
}

func NewGoogleService(db infrastructure.Database) GoogleService{
	return GoogleService{db}
}

func (gs GoogleService) CallbackHandler(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"callback",
	})
}