package routes

// import (
// 	"dependency_injection_tut/api/controller"
// 	"dependency_injection_tut/infrastructure"
// )

// type GoogleRoutes struct{
// 	server *infrastructure.Server
// 	routeController *controller.GoogleController
// }

// func NewGoogleRoute(ser *infrastructure.Server, cntrl *controller.GoogleController) *GoogleRoutes{
// 	return &GoogleRoutes{
// 		server: ser,
// 		routeController: cntrl,
// 	}
// }

// func (ggRoutes *GoogleRoutes) Setup(){
// 	api:= ggRoutes.server.Group("/api")

// 	api.GET("/google/callback",ggRoutes.routeController.CallbackHandler)

// }