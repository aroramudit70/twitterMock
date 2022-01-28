//go:build wireinject
// +build wireinject

package main

import (
	"context"
	controller "twitterMock/app"

	app2 "twitterMock/pkg/app"
	config2 "twitterMock/pkg/config"
	logger "twitterMock/pkg/logger"

	"github.com/google/wire"
)

func Initialize(ctx context.Context, envFile config2.Path) (*app2.App, func(), error) {
	panic(
		wire.Build(
			controller.HandlerModule,
			controller.DaoModulle,
			controller.ServiceModule,
			logger.NewConfig,
			logger.New,
			app2.Module,
		),
	)
}
