package bootstrap

import (
	"context"
	"go-fiber/api/controllers"
	"go-fiber/api/routes"
	"go-fiber/infrastructure"

	"go.uber.org/fx"
)

// Modules exported for initilization of the application
var Module = fx.Options(
	infrastructure.Module,
	routes.Module,
	controllers.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	handler infrastructure.Router,
	routes routes.Routes,
	logger infrastructure.Logger,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting the application")
			logger.Zap.Info("---------------------------")
			logger.Zap.Info("------- Go-fiber 🚀  -------")
			logger.Zap.Info("---------------------------")

			go func() {
				routes.Setup()
				handler.Listen(":5000")
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
