package api

import (
	"context"
	"go-weather/app/internal/models"

	"github.com/google/uuid"
)

type WeatherService interface {
	GetWeather(ctx context.Context, city string) (*models.CurrentWeather, error)
	CreateSubscription(context.Context, models.SubscribeRequest) error
	ConfirmSubscription(context.Context, uuid.UUID) error
	CancelSubscription(context.Context, uuid.UUID) error
}
