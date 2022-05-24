package routes

import "go.uber.org/fx"



var Module = fx.Options(
	fx.Provide(NewGoogleRoute),
	fx.Provide(NewRoutes),
)

type Routes []Route
type Route interface{
	Setup()
}

func NewRoutes(googleRoutes *GoogleRoutes) Routes{
	return Routes{
		googleRoutes,
	}
}

func (r Routes) Setup(){
	for _, route:=range r{
		route.Setup()
	}
}






