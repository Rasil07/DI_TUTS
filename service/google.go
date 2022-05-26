package service

// import (
// 	"dependency_injection_tut/infrastructure"
// 	"dependency_injection_tut/repository"

// 	"github.com/gin-gonic/gin"
// )

// type GoogleService struct{
// 	repository.Database
// }

// func NewGoogleService(db infrastructure.Database) GoogleService{
// 	return GoogleService{db}
// }

// func (gs GoogleService) CallbackHandler(c *gin.Context){
// 	c.JSON(200,gin.H{
// 		"message":"callback",
// 	})
// }