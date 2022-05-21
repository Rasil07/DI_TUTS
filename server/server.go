package server

import (
	"dependency_injection_tut/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct{
	engine *gin.Engine
	// Routes *routes.NewRoutes
}



func NewApp(s *gin.Engine) *Server{
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	
	ser:= Server{engine: s}
	 routes.NewRoute(ser.engine)
	
	return &ser
}

