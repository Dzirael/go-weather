package render

import (
	"bytes"
	"fmt"
	"go-weather/app/internal/models"
	"go-weather/templates"
	"text/template"
)

func ConfirmationMail(params ConfirmationEmailParams) (*models.MailMessage, error) {
	msg, err := render(templates.EmailConfirmSubscription, params)
	if err != nil {
		return nil, fmt.Errorf("apply template: %w", err)
	}

	return &models.MailMessage{
		Subject: "Conrirm subscription",
		Message: msg,
	}, nil
}

func WeatherUpdateMail(params WeatherUpdateParams) (*models.MailMessage, error) {
	msg, err := render(templates.EmailWeatherUpdate, params)
	if err != nil {
		return nil, fmt.Errorf("apply template: %w", err)
	}

	return &models.MailMessage{
		Subject: "Weather Notification",
		Message: msg,
	}, nil
}

func render(raw string, data any) (string, error) {
	tmpl, err := template.New("embedded").Parse(raw)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}

	return buf.String(), nil
}
