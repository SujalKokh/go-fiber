package bootstrap

import (
	"context"
	"go-fiber/infrastructure"

	"go.uber.org/fx"
)

// Modules exported for initilization of the application
var Module = fx.Options(
	infrastructure.Module,
	fx.Invoke(bootstrap),
)

func bootstrap(
	lifecycle fx.Lifecycle,
	logger infrastructure.Logger,
	handler infrastructure.Router,
) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Zap.Info("Starting the application")
			logger.Zap.Info("---------------------------")
			logger.Zap.Info("------- Go-fiber ♾️  -------")
			logger.Zap.Info("---------------------------")

			go func() {
				handler.Listen(":5000")
			}()
			return nil
		},
		OnStop: func(context.Context) error {
			return nil
		},
	})
}
