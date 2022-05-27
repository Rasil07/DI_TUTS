package routes

import (
	"dependency_injection_tut/api/controller"
	"dependency_injection_tut/api/middleware"
	"dependency_injection_tut/infrastructure"
)

type UserRoutes struct{
	server *infrastructure.Server
	routeController *controller.UserController
	userMiddleware *middleware.AuthMiddleware
}

func NewUserRoutes(srv *infrastructure.Server, ctrlr *controller.UserController, midl *middleware.AuthMiddleware) *UserRoutes{
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
	route.GET("/",ur.userMiddleware.Authorized(),ur.routeController.GetAll)
}