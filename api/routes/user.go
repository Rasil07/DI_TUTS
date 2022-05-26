package routes

import (
	"dependency_injection_tut/api/controller"
	"dependency_injection_tut/api/middleware"
	"dependency_injection_tut/infrastructure"
)

type UserRoutes struct{
	server *infrastructure.Server
	routeController *controller.UserController
	userMiddleware middleware.CreateUserPasswordMiddleware
}

func NewUserRoutes(srv *infrastructure.Server, ctrlr *controller.UserController, midl middleware.CreateUserPasswordMiddleware) *UserRoutes{
	return &UserRoutes{
		server: srv,
		routeController: ctrlr,
		userMiddleware: midl,
	}
}

func (ur *UserRoutes) Setup(){
	route := ur.server.Group("/api/users")

	route.POST("/",ur.routeController.Create)
	route.POST("login",ur.routeController.Login)
}