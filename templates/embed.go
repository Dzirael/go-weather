package templates

import (
	_ "embed"
)

//go:embed emails/confirm_subscription.gohtml
var EmailConfirmSubscription string

//go:embed emails/weather_update.gohtml
var EmailWeatherUpdate string
