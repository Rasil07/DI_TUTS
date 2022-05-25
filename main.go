package main

import (
	"context"
	"dependency_injection_tut/api/routes"
	"dependency_injection_tut/appRegiter"
	"dependency_injection_tut/infrastructure"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main(){
		godotenv.Load()
		app:= fx.New(
			appRegiter.Module,
			fx.Invoke(registerHooks),
		)

		app.Run()
}

func registerHooks(
	lifecycle fx.Lifecycle,
	ser *infrastructure.Server,
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