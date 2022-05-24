package appRegiter

import (
	"context"

	"go.uber.org/fx"
)


func Start(){
	ctx:=context.Background()
	
	app:= fx.New(
		Module,
	)

	app.Start(ctx)
	defer app.Stop(ctx)

	
	
}