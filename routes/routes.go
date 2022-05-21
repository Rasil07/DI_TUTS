package routes

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type NewRoutes struct{
	handler *gin.Engine
}


func NewRoute(h *gin.Engine) *NewRoutes{
	fmt.Println(h)
	r:= NewRoutes{handler: h}
	r.Setup()
	return &r
}

func (r *NewRoutes) Setup(){
	rout:=r.handler.Group("/api")

	rout.GET("/ping",r.Gethandler)
}

func (s *NewRoutes) Gethandler(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"pong",
	})
}






