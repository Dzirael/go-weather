package render

import "github.com/google/uuid"

type ConfirmationEmailParams struct {
	WebsiteURL string
	Token      uuid.UUID
}

type WeatherUpdateParams struct {
	WebsiteURL  string
	Token       uuid.UUID
	Temperature float64
	Humidity    float64
	Description string
}
