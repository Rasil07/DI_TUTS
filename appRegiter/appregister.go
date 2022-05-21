package appregiter

import (
	"dependency_injection_tut/server"

	"go.uber.org/fx"
)

var Module = fx.Options(
	// fx.Provide(routes.NewRoute),
	fx.Invoke(server.NewApp),
)