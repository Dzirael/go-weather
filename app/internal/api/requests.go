package api

import (
	"errors"
	"fmt"
	"go-weather/app/internal/models"
	"net/mail"
	"strings"
)

type SubscribeRequest struct {
	Email     string           `form:"email"`
	City      string           `form:"city"`
	Frequency models.Frequency `form:"frequency"`
}

func (r SubscribeRequest) Validate() error {
	if _, err := mail.ParseAddress(r.Email); err != nil {
		return errors.New("invalid email address")
	}

	if strings.TrimSpace(r.City) == "" {
		return errors.New("city is required")
	}

	if r.Frequency != models.Hourly && r.Frequency != models.Daily {
		return fmt.Errorf("frequency must be hourly or daily")
	}

	return nil
}
