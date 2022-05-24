package routes

import (
	"dependency_injection_tut/controller"
	"dependency_injection_tut/server"
)


type GoogleRoutes struct{
	server *server.Server
	routeController *controller.GoogleController
}

func NewGoogleRoute(ser *server.Server, cntrl *controller.GoogleController) *GoogleRoutes{
	return &GoogleRoutes{
		server: ser,
		routeController: cntrl,
	}
}

func (ggRoutes *GoogleRoutes) Setup(){
	api:= ggRoutes.server.Group("/api")

	api.GET("/google/callback",ggRoutes.routeController.CallbackHandler)

}