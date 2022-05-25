package appRegiter

import (
	"dependency_injection_tut/api/controller"
	"dependency_injection_tut/api/routes"
	"dependency_injection_tut/infrastructure"
	"dependency_injection_tut/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	routes.Module,
	controller.Module,
	infrastructure.Module,
	service.Module,
)