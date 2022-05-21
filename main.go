package main

import (
	"context"
	"dependency_injection_tut/server"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

func main(){
		
		app:= fx.New(
			fx.Provide(gin.Default),
			server.Module,
			// fx.Provide(routes.NewRoute),
			fx.Invoke(registerHooks),
		)

		app.Run()
}

func registerHooks(
	lifecycle fx.Lifecycle,
	ser *gin.Engine,
){
	lifecycle.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go ser.Run()
				return nil
			},
		},
	)
}