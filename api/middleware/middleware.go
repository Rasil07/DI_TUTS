package middleware

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(NewCreateUserPasswordMiddleware),
	fx.Provide(NewAuthMidleware),
)