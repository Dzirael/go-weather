package api

import (
	"errors"
	"fmt"
	"go-weather/app/internal/api/resp"
	"go-weather/app/internal/apperrors"
	"go-weather/app/internal/config"
	"go-weather/app/internal/models"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

type Handler struct {
	weatherService WeatherService
	config         *config.Config
}

func NewHandler(weatherService WeatherService) *Handler {
	return &Handler{weatherService: weatherService}
}

func (h *Handler) GetWeather(c fiber.Ctx) error {
	city := c.Query("city")
	if city == "" {
		return resp.NotFound(fmt.Errorf("invalid request"), resp.ErrZeroCode, "parse query").Respond(c)
	}

	result, err := h.weatherService.GetWeather(c.Context(), city)
	if err != nil {
		if errors.Is(err, apperrors.ErrCityNotSupported) {
			return resp.NotFound(err, resp.ErrZeroCode, "city not found").Respond(c)
		}
		return fmt.Errorf("get user: %w", err)
	}

	return c.JSON(result)
}

func (h *Handler) CreateSubscription(c fiber.Ctx) error {
	var req SubscribeRequest
	if err := c.Bind().Form(&req); err != nil {
		return resp.BadRequest(err, resp.ErrZeroCode, "parse form").Respond(c)
	}

	if err := req.Validate(); err != nil {
		return resp.BadRequest(err, resp.ErrZeroCode, "validate form").Respond(c)
	}

	err := h.weatherService.CreateSubscription(c.Context(), models.SubscribeRequest{
		Email:     req.Email,
		City:      req.City,
		Frequency: req.Frequency,
	})
	if err != nil {
		if errors.Is(err, apperrors.ErrAlreadyHaveSubscription) {
			return resp.Conflict(err, resp.ErrZeroCode, "email already subscribed").Respond(c)
		}
		return fmt.Errorf("create subscription: %w", err)
	}

	return nil
}

func (h *Handler) ConfirmSubscription(c fiber.Ctx) error {
	if c.Params("token") == "" {
		return resp.NotFound(errors.New("token not found"), resp.ErrZeroCode, "token not found").Respond(c)
	}

	token, err := uuid.Parse(c.Params("token"))
	if err != nil {
		return resp.BadRequest(err, resp.ErrZeroCode, "parse token").Respond(c)
	}

	err = h.weatherService.ConfirmSubscription(c.Context(), token)
	if err != nil {
		return fmt.Errorf("confirm subscription: %w", err)
	}

	return nil
}

func (h *Handler) CancelSubscription(c fiber.Ctx) error {
	if c.Params("token") == "" {
		return resp.NotFound(errors.New("token not found"), resp.ErrZeroCode, "token not found").Respond(c)
	}

	token, err := uuid.Parse(c.Params("token"))
	if err != nil {
		return resp.BadRequest(err, resp.ErrZeroCode, "parse token").Respond(c)
	}

	err = h.weatherService.CancelSubscription(c.Context(), token)
	if err != nil {
		return fmt.Errorf("cancel subscription: %w", err)
	}

	return nil
}
