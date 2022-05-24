package controller

import (
	"dependency_injection_tut/service"

	"github.com/gin-gonic/gin"
)

type GoogleController struct{
	googleService service.GoogleService
}

func NewGoogleController(gs service.GoogleService) *GoogleController{
	return &GoogleController{googleService: gs}
}

func (cnt *GoogleController) CallbackHandler(c *gin.Context){
 cnt.googleService.CallbackHandler(c)
}