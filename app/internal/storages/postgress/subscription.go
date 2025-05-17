package repository

import (
	"context"
	"errors"
	"fmt"
	"go-weather/app/internal/apperrors"
	"go-weather/app/internal/models"
	"go-weather/app/internal/storages/postgress/sqlc"

	"github.com/google/uuid"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type WeatherRepository struct {
	repo *sqlc.Queries
	pool *pgxpool.Pool
}

func NewWeatherRepository(repo *sqlc.Queries, pool *pgxpool.Pool) *WeatherRepository {
	return &WeatherRepository{repo: repo, pool: pool}
}

func (r WeatherRepository) CreateSubscription(ctx context.Context, params models.CreateSubscriptionParams) error {
	err := r.repo.CreateSubscription(ctx, sqlc.CreateSubscriptionParams{
		ID:        params.ID,
		Code:      params.ConfirmationCode,
		Status:    string(params.Status),
		Frequency: string(params.Frequency),
		Email:     params.Email,
		City:      params.City,
	})
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				return apperrors.ErrAlreadyHaveSubscription
			}
		}
		return fmt.Errorf("sqlc: create subscription: %w", err)
	}
	return nil
}

func (r WeatherRepository) UpdateSubscriptionStatusByCode(ctx context.Context, status models.SubscriptionStatus, confirmationCode uuid.UUID) error {
	err := r.repo.UpdateSubscriptionStatusByCode(ctx, string(status), confirmationCode)
	if errors.Is(err, pgx.ErrNoRows) {
		return apperrors.ErrConfirmationCodeNotFound
	}

	if err != nil {
		return fmt.Errorf("sqlc: failed to update subscription status: %w", err)
	}

	return nil
}

func (r WeatherRepository) SetSendetNow(ctx context.Context, subscriptionID uuid.UUID) error {
	return r.repo.SetSendedNow(ctx, subscriptionID)
}

func (r WeatherRepository) GetWaitingsSubscription(ctx context.Context) ([]uuid.UUID, error) {
	subscriptionIDs, err := r.repo.GetWaitingsNotification(ctx)
	if err != nil {
		return nil, fmt.Errorf("sqlc: get waiting notification: %w", err)
	}

	return subscriptionIDs, nil
}

func (r WeatherRepository) GetSubscriptionByID(ctx context.Context, subscriptionID uuid.UUID) (*models.Subscription, error) {
	sub, err := r.repo.GetSubscriptionByID(ctx, subscriptionID)
	if err != nil {
		return nil, fmt.Errorf("sqlc: get subscription: %w", err)
	}

	return subscriptionToModel(sub), nil
}

func subscriptionToModel(sub sqlc.Subscription) *models.Subscription {
	return &models.Subscription{
		ID:               sub.ID,
		ConfirmationCode: &sub.ConfirmationCode,
		Status:           models.SubscriptionStatus(sub.Status),
		Frequency:        models.Frequency(sub.Frequency),
		Email:            sub.Email,
		City:             sub.City,
	}

}

func (r WeatherRepository) DeleteSubscription(ctx context.Context, code uuid.UUID) error {
	return r.repo.DeleteSubscription(ctx, code)
}
