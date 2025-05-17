package service

import (
	"context"
	"fmt"
	"go-weather/app/internal/apperrors"
	"go-weather/app/internal/config"
	"go-weather/app/internal/models"
	"go-weather/app/internal/render"

	"github.com/google/uuid"
)

type Email interface {
	Send(to string, msg *models.MailMessage) error
}

type WeatherClient interface {
	GetWeather(ctx context.Context, sity string) (*models.CurrentWeather, error)
}

type WeatherService struct {
	weatherClient WeatherClient
	repo          WeatherRepository
	smtp          Email
	config        *config.Config
}

func NewWeatherService(weatherClient WeatherClient, repo WeatherRepository, smtp Email, cfg *config.Config) *WeatherService {
	return &WeatherService{weatherClient: weatherClient, repo: repo, smtp: smtp, config: cfg}
}

func (s WeatherService) GetWeather(ctx context.Context, city string) (*models.CurrentWeather, error) {
	return s.weatherClient.GetWeather(ctx, city)
}

func (s WeatherService) CreateSubscription(ctx context.Context, params models.SubscribeRequest) error {
	// Need add verification that city is supported
	subscriptionID := uuid.Must(uuid.NewV7())
	confirmationCode := uuid.Must(uuid.NewV7())

	have, err := s.repo.HaveActiveSubscription(ctx, params.Email)
	if err != nil {
		return fmt.Errorf("check email: %w", err)
	}

	if have {
		return apperrors.ErrAlreadyHaveSubscription
	}

	err = s.repo.CreateSubscription(ctx, models.CreateSubscriptionParams{
		ID:               subscriptionID,
		ConfirmationCode: confirmationCode,
		Status:           models.SubscriptionPending,
		Frequency:        params.Frequency,
		Email:            params.Email,
		City:             params.City,
	})
	if err != nil {
		return fmt.Errorf("save subscription: %w", err)
	}

	confirmMail, err := render.ConfirmationMail(render.ConfirmationEmailParams{
		WebsiteURL: s.config.WebsiteURL,
		Token:      confirmationCode,
	})
	if err != nil {
		return fmt.Errorf("format mail: %w", err)
	}

	return s.smtp.Send(params.Email, confirmMail)
}

func (s WeatherService) ConfirmSubscription(context.Context, uuid.UUID) error {
	return nil
}
func (s WeatherService) CancelSubscription(context.Context, uuid.UUID) error {
	return nil
}

func (s WeatherService) HandleTriggeredNotification(ctx context.Context, subscriptionID uuid.UUID) error {
	subscription, err := s.repo.GetSubscriptionByID(ctx, subscriptionID)
	if err != nil {
		return fmt.Errorf("get subscription: %w", err)
	}

	currentWeather, err := s.weatherClient.GetWeather(ctx, subscription.City)
	if err != nil {
		return fmt.Errorf("get weather: %w", err)
	}

	nofiticationMail, err := render.WeatherUpdateMail(render.WeatherUpdateParams{
		WebsiteURL:  s.config.WebsiteURL,
		Token:       *subscription.ConfirmationCode,
		Temperature: currentWeather.Temperature,
		Humidity:    currentWeather.Humidity,
		Description: currentWeather.Description,
	})
	if err != nil {
		return fmt.Errorf("format mail: %w", err)
	}

	err = s.smtp.Send(subscription.Email, nofiticationMail)
	if err != nil {
		return fmt.Errorf("send mail: %w", err)
	}

	return s.repo.SetSendetNow(ctx, subscriptionID)
}
