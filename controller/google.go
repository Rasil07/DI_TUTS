package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type GoogleController struct{}

func NewGoogleController() *GoogleController{

	return &GoogleController{}
}

func (cnt *GoogleController) CallbackHandler(c *gin.Context){
fmt.Println("Callback handler reached")
c.JSON(200,gin.H{
"message":"Google callback",
})
}