package models

import (
	"github.com/google/uuid"
)

type SubscribeRequest struct {
	Email     string    `json:"email" form:"email"`
	City      string    `json:"city" form:"city"`
	Frequency Frequency `json:"frequency" form:"frequency"`
}

// DTO
type CreateSubscriptionParams struct {
	ID               uuid.UUID
	ConfirmationCode uuid.UUID
	Status           SubscriptionStatus
	Frequency        Frequency
	Email            string
	City             string
}

type CurrentWeather struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
	Description string  `json:"description"`
}

const NoWeatherDescription = "no description"

type SubscriptionStatus string

const (
	SubscriptionActive   SubscriptionStatus = "ACTIVE"
	SubscriptionCanceled SubscriptionStatus = "CANCELED"
	SubscriptionPending  SubscriptionStatus = "PENDING"
)

type Frequency string

const (
	Hourly Frequency = "hourly"
	Daily  Frequency = "daily"
)

type Subscription struct {
	ID               uuid.UUID
	ConfirmationCode *uuid.UUID
	Status           SubscriptionStatus
	Frequency        Frequency
	Email            string
	City             string
}
