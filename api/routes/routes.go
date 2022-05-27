package routes

import "go.uber.org/fx"



var Module = fx.Options(
	fx.Provide(NewUserRoutes),
	fx.Provide(NewImageRoutes),
	fx.Provide(NewRoutes),
)

type Routes []Route
type Route interface{
	Setup()
}

func NewRoutes(userRoutes *UserRoutes,imageRoutes *ImageRoutes) Routes{
	return Routes{
		userRoutes,
		imageRoutes,
	}
}

func (r Routes) Setup(){
	for _, route:=range r{
		route.Setup()
	}
}






