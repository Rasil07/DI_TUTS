package server

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)


var Module = fx.Options(
	fx.Provide(NewApp),
)
type Server struct{
	 *gin.Engine
}

func NewApp(s *gin.Engine) *Server{
	s.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))
	ser:= Server{s}
	return &ser
}

