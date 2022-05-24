package service

import (
	"dependency_injection_tut/database"

	"github.com/gin-gonic/gin"
)

type GoogleService struct{
	database.Database
}

func NewGoogleService(db database.Database) GoogleService{
	return GoogleService{db}
}

func (gs GoogleService) CallbackHandler(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"callback",
	})
}