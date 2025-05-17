package service

import (
	"context"
	"go-weather/app/internal/models"

	"github.com/google/uuid"
)

type WeatherRepository interface {
	CreateSubscription(ctx context.Context, params models.CreateSubscriptionParams) error
	UpdateSubscriptionStatusByCode(ctx context.Context, status models.SubscriptionStatus, confirmationCode uuid.UUID) error
	GetSubscriptionByID(ctx context.Context, subscriptionID uuid.UUID) (*models.Subscription, error)
	GetWaitingsSubscription(ctx context.Context) ([]uuid.UUID, error)
	SetSendetNow(ctx context.Context, subscriptionID uuid.UUID) error
	DeleteSubscription(ctx context.Context, code uuid.UUID) error
}
