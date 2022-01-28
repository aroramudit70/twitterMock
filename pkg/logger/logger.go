package logger

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	Enabled            bool   `envconfig:"LOG_ENABLED"`
	Level              string `envconfig:"LOG_LEVEL"`
	TimestampFormat    string `envconfig:"LOG_TIMESTAMP_FORMAT"`
	StdoutEnabled      bool   `envconfig:"LOG_STDOUT_ENABLED"`
	FilePath           string `envconfig:"LOG_FILE"`
	MaxSize            int    `envconfig:"LOG_FILE_MAX_SIZE"`
	MaxBackups         int    `envconfig:"LOG_FILE_MAX_BACKUPS"`
	MaxAge             int    `envconfig:"LOG_FILE_MAX_AGE"`
	CompressionEnabled bool   `envconfig:"LOG_COMPRESSION_ENABLED"`
	StacktraceEnabled  bool   `envconfig:"LOG_STACKTRACE_ENABLED"`
	TrueIP             string `envconfig:"LOG_TRUE_IP"`
}

func Default() *zap.Logger {
	zapConfig := zap.NewProductionConfig()
	zapConfig.OutputPaths = []string{"stdout"}
	zapConfig.DisableStacktrace = true
	zapConfig.Sampling = nil

	logger, _ := zapConfig.Build()
	logger.Sync()
	zapConfig.Level.UnmarshalText([]byte("debug"))
	return logger
}

type lumberjackSink struct {
	*lumberjack.Logger
}

// Sync implements zap.Sink
func (lumberjackSink) Sync() error { return nil }

func NewConfig(config *Config) (*zap.Config, error) {
	zapConfig := zap.NewProductionConfig()
	zapConfig.Sampling = nil
	zapConfig.DisableStacktrace = !config.StacktraceEnabled

	zapConfig.OutputPaths = []string{fmt.Sprintf("lumberjack:%s", config.FilePath)}
	if config.StdoutEnabled {
		zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stdout")
	}
	if config.TimestampFormat != "" {
		encoderCfg := zap.NewProductionEncoderConfig()
		switch config.TimestampFormat {
		case "RFC3339":
			encoderCfg.EncodeTime = zapcore.RFC3339TimeEncoder
		case "EpochMillis":
			encoderCfg.EncodeTime = EpochMillisTimeEncoder
		case "CustomVNSP":
			encoderCfg.EncodeTime = CustomVNSPTimeEncoder
		}
		encoderCfg.TimeKey = "log_time_stamp"
		encoderCfg.LevelKey = "log_level"
		encoderCfg.MessageKey = "log_message"
		encoderCfg.CallerKey = "function_name"
		encoderCfg.EncodeLevel = zapcore.CapitalLevelEncoder
		zapConfig.EncoderConfig = encoderCfg
	}
	if config.Level != "" {
		if err := zapConfig.Level.UnmarshalText([]byte(strings.ToLower(config.Level))); err != nil {
			return nil, errors.Errorf("Unable to parse log level from config %s: %s", config.Level, err.Error())
		}
	}
	return &zapConfig, nil
}

func New(config *Config, zapConfig *zap.Config) (*zap.Logger, func(), error) {
	if !config.Enabled {
		return zap.NewNop(), func() {}, nil
	}

	err := zap.RegisterSink("lumberjack", func(u *url.URL) (zap.Sink, error) {
		return lumberjackSink{
			Logger: &lumberjack.Logger{
				Filename:   config.FilePath,
				MaxSize:    config.MaxSize,
				MaxBackups: config.MaxBackups,
				MaxAge:     config.MaxAge,
				Compress:   config.CompressionEnabled,
			},
		}, nil
	})
	if err != nil {
		return nil, nil, err
	}

	logger, err := zapConfig.Build()
	if err != nil {
		return nil, nil, err
	}

	// hostname, err := os.Hostname()
	// if err != nil {
	// 	return nil, nil, err
	// }

	logger = logger.With(
	// zap.String("VSAD_ID", "GVNV"),
	// zap.String("app_name", "EDS_CLIENT_APP"),
	// zap.String("app_type", "MEC"),
	// zap.String("app_version", build2.Release),
	// zap.String("vast_id", "25355"),
	// zap.String("version", "1"),
	// zap.String("type", "System"),
	// zap.String("server_host", hostname),
	// zap.String("server_port", "13496"),
	// zap.String("True_ip", config.TrueIP),
	// zap.String("app_environment", config.TrueIP),
	)
	stop := func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Error("Panic while syncing zap logger")
			}
		}()
		if logger != nil {
			logger.Sync()
		}
	}
	return logger, stop, nil
}

func EpochMillisTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	nanos := t.UnixNano()
	millis := nanos / int64(time.Millisecond)
	enc.AppendInt64(millis)
}

func CustomVNSPTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.UTC().Format("2006-01-02 15:04:05"))
}
