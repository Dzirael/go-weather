package notification

import (
	"context"
	"fmt"
	"go-weather/app/internal/config"
	"go-weather/app/internal/service"
	"go-weather/app/pkg/jobrunner"
	"go-weather/app/pkg/logger"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

type NotificationTracker struct {
	repo    service.WeatherRepository
	service *service.WeatherService
}

func Run(lc fx.Lifecycle,
	cfg *config.Config, repo service.WeatherRepository, service *service.WeatherService,
) {
	tracker := &NotificationTracker{repo, service}

	interval := cfg.BackgroundJobs.WeatherSubscriptionCheckInterval

	job := jobrunner.New(logger.L.Logger,
		"Track subscription start",
		interval,
		func() {

			logger.L.Info("start checkCurrenSubscription()")

			err := tracker.checkCurrenSubscription()
			if err != nil {
				logger.L.Error("check expired orders", zap.Error(err), zap.String("service", "expire_tracker"))
			}
		},
	)

	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			job.Start()
			return nil
		},
		OnStop: func(_ context.Context) error {
			job.Stop()
			return nil
		},
	})
}

func (t *NotificationTracker) checkCurrenSubscription() error {
	ctx := context.Background()

	subscriptionIDs, err := t.repo.GetWaitingsSubscription(ctx)
	if err != nil {
		return fmt.Errorf("get expired orders: %w", err)
	}

	logger.L.Info("got current notifications", zap.Any("ids", subscriptionIDs))

	for _, id := range subscriptionIDs {
		err = t.service.HandleTriggeredNotification(ctx, id)
		if err != nil {
			logger.L.Error("handle triggered subscription",
				zap.Error(err),
				zap.String("sub_id", id.String()),
			)
		}
	}

	return nil
}
