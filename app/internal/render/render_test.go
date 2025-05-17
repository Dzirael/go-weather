package render_test

import (
	"go-weather/app/internal/render"
	"strings"
	"testing"

	"github.com/google/uuid"
)

func TestConfirmationMail(t *testing.T) {
	params := render.ConfirmationEmailParams{
		Token:      uuid.Must(uuid.NewV7()),
		WebsiteURL: "http://mysite.com",
	}

	email, err := render.ConfirmationMail(params)
	if err != nil {
		t.Fatalf("ConfirmationMail() error = %v", err)
	}

	if !strings.Contains(email.Message, params.Token.String()) {
		t.Errorf("email message does not contain token: %s", params.Token)
	}

	if email.Subject != "Conrirm subscription" {
		t.Errorf("unexpected topic: %s", email.Subject)
	}
}

func TestWeatherUpdateMail(t *testing.T) {
	params := render.WeatherUpdateParams{
		Temperature: 22.5,
		Humidity:    55,
		Description: "Clear sky",
		Token:       uuid.Must(uuid.NewV7()),
		WebsiteURL:  "http://mysite.com/bye",
	}

	email, err := render.WeatherUpdateMail(params)
	if err != nil {
		t.Fatalf("WeatherUpdateMail() error = %v", err)
	}

	if !strings.Contains(email.Message, params.Description) {
		t.Errorf("email message does not contain description: %s", params.Description)
	}

	if !strings.Contains(email.Message, params.Token.String()) {
		t.Errorf("email message does not contain token: %s", params.Token)
	}

	if email.Subject != "Weather Notification" {
		t.Errorf("unexpected topic: %s", email.Subject)
	}
}
