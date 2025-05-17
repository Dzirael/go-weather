package api

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"go-weather/app/internal/api/middleware"
	"go-weather/app/internal/config"
	"go-weather/app/pkg/logger"
	"go-weather/docs"
)

type Params struct {
	fx.In

	Server  *fiber.App `name:"public-api"`
	Config  *config.Config
	Logger  *logger.Logger
	Handler *Handler
}

type Result struct {
	fx.Out
	Server *fiber.App `name:"public-api"`
}

func setupDocs(app *fiber.App, cfg *config.Config) {
	app.Group("/swagger").
		Get("/", func(c fiber.Ctx) error {
			const link = "https://validator.swagger.io/?url=%s/swagger/doc.yaml"
			return c.Redirect().To(fmt.Sprintf(link, cfg.ApplicationURL))
		}).
		Get("/doc.yaml", func(c fiber.Ctx) error {
			return c.SendString(docs.SwaggerSpecYaml)
		})
}

func RegisterRoutes(s Params) {
	app := s.Server
	setupDocs(app, s.Config)

	app.Get("/health", func(ctx fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"status": "ok"})
	})

	app.Get("/weather", s.Handler.GetWeather)

	app.Post("/subscribe", s.Handler.CreateSubscription)
	app.Get("/confirm/:token", s.Handler.ConfirmSubscription)
	app.Get("/unsubscribe/:token", s.Handler.CancelSubscription)
}

func NewServer(lc fx.Lifecycle, cfg *config.Config, l *logger.Logger) (Result, error) {
	fiberConfig := fiber.Config{
		ErrorHandler: middleware.ErrorHandler(l),
	}

	if cfg.Server.ProxyHeader != "" {
		l.Info("fiber: using ProxyHeader from config",
			zap.String("proxy_header", cfg.Server.ProxyHeader),
		)
		fiberConfig.ProxyHeader = cfg.Server.ProxyHeader
	}

	app := fiber.New(fiberConfig)

	app.Use(cors.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))
	app.Use(middleware.RequestID())
	// app.Use(middleware.LoggerMiddleware(l))
	// app.Use(middleware.LoggerRequestMiddleware(l))

	l.Info("server started", zap.String("listening", cfg.Server.Address))

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				if err := app.Listen(cfg.Server.Address); err != nil {
					l.Error("failed to start http server", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			return app.Shutdown()
		},
	})

	return Result{Server: app}, nil
}
