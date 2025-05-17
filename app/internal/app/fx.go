package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"

	"go-weather/app/internal/api"
	"go-weather/app/internal/clients/smtp"
	"go-weather/app/internal/clients/wttr"
	"go-weather/app/internal/config"
	"go-weather/app/internal/service"
	"go-weather/app/internal/service/cron/notification"
	repository "go-weather/app/internal/storages/postgress"
	"go-weather/app/pkg/logger"
)

func Run() {
	log := logger.New()
	defer log.Sync()

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	options := []fx.Option{
		fx.Supply(log),
		fx.Supply(ctx),

		fx.WithLogger(func() fxevent.Logger {
			return fxevent.NopLogger
		}),

		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		//	   return &fxevent.ZapLogger{Logger: log}
		// }),

		fx.Provide(config.New),

		fx.Provide(repository.NewPostgres),
		fx.Provide(fx.Annotate(
			repository.NewWeatherRepository,
			fx.As(new(service.WeatherRepository)),
		)),

		fx.Provide(fx.Annotate(
			smtp.New,
			fx.As(new(service.Email)),
		)),
		fx.Provide(fx.Annotate(
			wttr.New,
			fx.As(new(service.WeatherClient)),
		)),

		fx.Provide(
			service.NewWeatherService,
			fx.Annotate(
				service.NewWeatherService,
				fx.As(new(api.WeatherService)),
			),
		),

		fx.Provide(
			api.NewHandler,
			api.NewServer,
		),
		fx.Invoke(api.RegisterRoutes),
		fx.Invoke(notification.Run),
	}

	if err := fx.ValidateApp(options...); err != nil {
		log.Fatal("failed to validate fx app", zap.Error(err))
	}

	app := fx.New(options...)
	if err := app.Start(ctx); err != nil {
		log.Fatal("failed to start app", zap.Error(err))
	}

	<-ctx.Done()

	if err := app.Stop(context.Background()); err != nil {
		log.Warn("failed to stop app", zap.Error(err))
	}
}
