package main

import (
	"context"
	"dependency_injection_tut/appRegiter"
	"dependency_injection_tut/routes"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main(){
		
		app:= fx.New(
			fx.Provide(gin.Default),
			appRegiter.Module,
			fx.Invoke(registerHooks),
		)

		app.Run()
}

func registerHooks(
	lifecycle fx.Lifecycle,
	ser *gin.Engine,
	rts routes.Routes,
){
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				rts.Setup()
				go ser.Run()
				return nil
			},
		},
	)
}