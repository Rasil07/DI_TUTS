package service

import "go.uber.org/fx"

var Module = fx.Options(
	// fx.Provide(NewGoogleService),
	fx.Provide(NewUserService),
	fx.Provide(NewImageService),

)