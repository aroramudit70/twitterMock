package main

import (
	"context"
	"flag"
	"os"

	"twitterMock/pkg/app"
	"twitterMock/pkg/config"

	"go.uber.org/zap"
)

var app1 *app.App

func main() {
	var env string
	flag.StringVar(&env, "config", "", ".env file path")
	flag.Parse()
	envPath := config.Path(env)
	app1, stop, err := Initialize(context.Background(), envPath)
	if err != nil {
		os.Exit(-1)
	}

	app1.AwaitShutdown(stop)
}

func StopApp(logger *zap.Logger) {
	stop := func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic while shutting down HTTP server")
			}
		}()
		logger.Sugar().Info("Stopping HTTP server.")
	}
	app1.AwaitShutdown(stop)
}
