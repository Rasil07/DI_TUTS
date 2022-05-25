package infrastructure

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)



type Server struct{
	 *gin.Engine
}

func NewApp() *Server{
	s := gin.Default()
	
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	ser:= Server{s}
	return &ser
}

