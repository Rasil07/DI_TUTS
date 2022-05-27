package routes

import (
	"dependency_injection_tut/api/controller"
	"dependency_injection_tut/infrastructure"
)


type ImageRoutes struct{
	server *infrastructure.Server
	routeHandler *controller.ImageController
}

func NewImageRoutes(ser *infrastructure.Server, ctrlr *controller.ImageController) *ImageRoutes{
	return &ImageRoutes{
		server: ser,
		routeHandler: ctrlr,
	}
}

func (imr *ImageRoutes) Setup(){
	route:= imr.server.Group("/api/image")
	imr.server.MaxMultipartMemory = 8 << 20
	route.POST("/resize", imr.routeHandler.HandleResize)
	
}