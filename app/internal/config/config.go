package config

import (
	"fmt"
	"time"

	"github.com/caarlos0/env/v10"
)

type Config struct {
	Server struct {
		Address     string `env:"SERVER_ADDRESS,required"`
		ProxyHeader string `env:"SERVER_PROXY_HEADER,required"`
	}

	ApplicationURL string `env:"APPLICATION_URL,required"`
	WebsiteURL     string `env:"WEBSITE_URL,required"`

	DisableStackTrace bool `env:"DISABLE_STACK_TRACE"`

	DB struct {
		PostgresDSN string `env:"POSTGRES_DSN,required"`
	}

	BackgroundJobs struct {
		WeatherSubscriptionCheckInterval time.Duration `env:"WEATHER_SUBSCRIPTION_CHECK_INTERVAL,required"`
	}

	SMTP struct {
		Host     string `env:"SMTP_HOST,required"`
		Port     int    `env:"SMTP_PORT,required"`
		Username string `env:"SMTP_USERNAME,required"`
		Password string `env:"SMTP_PASSWORD,required"`
	}
}

func New() (*Config, error) {
	var cfg Config
	err := env.Parse(&cfg)
	if err != nil {
		return nil, fmt.Errorf("parse env: %w", err)
	}

	return &cfg, nil
}
