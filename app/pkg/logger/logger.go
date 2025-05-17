package logger

import (
	"os"

	"go.uber.org/zap"
)

type Environment string

const (
	Dev   Environment = "dev"
	Stage Environment = "stage"
	Prod  Environment = "prod"
)

var L *Logger = &Logger{
	Logger:      zap.Must(zap.NewDevelopment()),
	Environment: Dev,
}

type Logger struct {
	*zap.Logger
	Environment Environment
}

func New() *Logger {
	var (
		disableStacktrace = os.Getenv("LOGGER_DISABLE_STACKTRACE") == "true"
		environment       = Environment(os.Getenv("ENVIRONMENT"))
	)

	cfg := zap.NewProductionConfig()
	if environment == Dev {
		cfg = zap.NewDevelopmentConfig()
	}

	cfg.DisableStacktrace = disableStacktrace
	l, err := cfg.Build()
	if err != nil {
		panic(err)
	}

	out := &Logger{
		Logger:      l,
		Environment: environment,
	}
	L = out
	return out
}
