package app

import (
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"twitterMock/pkg/config"
	"twitterMock/pkg/httpserver"
	"twitterMock/pkg/logger"
	"twitterMock/pkg/mongodb"

	"github.com/google/wire"

	"go.uber.org/zap"
)

type (
	Config struct {
		Logger     *logger.Config
		HTTPServer *httpserver.Config
		MongoDB    mongodb.Config
	}

	App struct {
		signals    chan os.Signal
		logger     *zap.Logger
		httpserver *http.Server
	}
)

var ConfigModule = wire.NewSet(
	NewConfig,
	wire.FieldsOf(
		new(*Config),
		"Logger",
		"HTTPServer",
		"MongoDB",
	),
)

func NewConfig(path config.Path) (*Config, error) {
	cfg := &Config{}
	if err := config.Load(path, cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func NewApp(
	logger *zap.Logger,
	httpserver *http.Server,
) *App {
	app := &App{
		logger:     logger,
		httpserver: httpserver,
	}

	lock := sync.RWMutex{}
	lock.Lock()
	app.signals = make(chan os.Signal, 1)
	signal.Notify(app.signals, syscall.SIGINT, syscall.SIGTERM)
	lock.Unlock()

	return app
}

func (a *App) AwaitShutdown(stop func()) {
	signal := <-a.signals
	a.logger.Info("Received shutdown signal", zap.String("signal", signal.String()))
	stop()
	os.Exit(0)
}

var Module = wire.NewSet(
	ConfigModule,
	httpserver.New,
	mongodb.GetConnection,
	NewApp,
)
