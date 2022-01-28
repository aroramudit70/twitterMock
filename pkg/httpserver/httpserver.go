package httpserver

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Config struct {
	Enabled         bool   `envconfig:"HTTP_SERVER_ENABLED"`
	Host            string `envconfig:"HTTP_SERVER_HOST"`
	Port            int    `envconfig:"HTTP_SERVER_PORT"`
	ListenLimit     int    `envconfig:"HTTP_SERVER_LISTEN_LIMIT"`
	KeepAlive       int    `envconfig:"HTTP_SERVER_KEEP_ALIVE"`
	ReadTimeout     int    `envconfig:"HTTP_SERVER_READ_TIMEOUT"`
	WriteTimeout    int    `envconfig:"HTTP_SERVER_WRITE_TIMEOUT"`
	ShutdownTimeout int    `envconfig:"HTTP_SERVER_SHUTDOWN_TIMEOUT"`
}

func New(ctx context.Context, config *Config, logger *zap.Logger, handler http.Handler) (*http.Server, func(), error) {
	if config.Enabled {
		address := fmt.Sprintf("%s:%d", config.Host, config.Port)
		server := &http.Server{
			Addr:         address,
			Handler:      handler,
			ReadTimeout:  time.Duration(config.ReadTimeout) * time.Second,
			WriteTimeout: time.Duration(config.WriteTimeout) * time.Second,
		}
		go func() {
			logger.Info("Starting HTTP server", zap.String("address", address))
			if err := server.ListenAndServe(); err != http.ErrServerClosed {
				logger.Sugar().Errorf("Failed to run HTTP server at address %s: %v", address, err)
			}
		}()
		stop := func() {
			defer func() {
				if err := recover(); err != nil {
					logger.Error("Panic while shutting down HTTP server")
				}
			}()
			logger.Sugar().Info("Stopping HTTP server.")
			context, cancel := context.WithTimeout(context.Background(), time.Duration(config.ShutdownTimeout)*time.Second)
			defer cancel()
			err := server.Shutdown(context)
			if err != nil {
				logger.Sugar().Errorf("Error while shutting down HTTP server", err.Error())
			}
		}
		return server, stop, nil
	}
	return nil, func() {}, nil
}
