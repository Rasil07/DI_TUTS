package appRegiter

import (
	"dependency_injection_tut/controller"
	"dependency_injection_tut/database"
	"dependency_injection_tut/routes"
	"dependency_injection_tut/server"
	"dependency_injection_tut/service"

	"go.uber.org/fx"
)

var Module = fx.Options(
	routes.Module,
	controller.Module,
	server.Module,
	database.Module,
	service.Module,
)