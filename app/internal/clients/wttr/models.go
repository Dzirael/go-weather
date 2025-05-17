package wttr

import "github.com/shopspring/decimal"

type WeatherResponse struct {
	CurrentCondition []CurrentWeather `json:"current_condition"`
}

type CurrentWeather struct {
	Humidity    decimal.Decimal `json:"humidity"`
	TempC       decimal.Decimal `json:"temp_C"`
	TempF       decimal.Decimal `json:"temp_F"`
	WeatherDesc []Value         `json:"weatherDesc"`
}

type Value struct {
	Value string `json:"value"`
}
