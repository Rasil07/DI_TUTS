package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Server struct{
	*gin.Engine
}

var Module = fx.Options(
	fx.Invoke(NewApp),
)

func NewApp(s *gin.Engine) *Server{
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	
	ser:= Server{s}
	ser.registerroutes()
	return &ser
}

func (s *Server) registerroutes(){
	route:=s.Group("/api")
	route.GET("/ping",s.Gethandler)
}

func (s *Server) Gethandler(c *gin.Context){
	c.JSON(200,gin.H{
		"message":"pong",
	})
}
