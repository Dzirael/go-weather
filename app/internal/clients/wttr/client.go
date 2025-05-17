package wttr

import (
	"context"
	"fmt"
	"go-weather/app/internal/models"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	client *resty.Client
}

const baseURL = "https://wttr.in"

func New() *Client {
	client := resty.New().SetBaseURL(baseURL)
	return &Client{
		client: client,
	}
}

func (c *Client) GetWeather(ctx context.Context, sity string) (*models.CurrentWeather, error) {
	var result WeatherResponse

	req := c.client.R().
		SetContext(ctx).
		SetQueryParam("format", "j1").
		SetResult(&result)

	resp, err := req.Get(sity)
	if err != nil {
		return nil, err
	}

	if resp.IsError() {
		return nil, fmt.Errorf(
			"request failed: status=%d, body=%s",
			resp.StatusCode(), resp.String(),
		)
	}

	if len(result.CurrentCondition) == 0 {
		return nil, fmt.Errorf("empty current_condition in response")
	}

	return weatherToModel(&result), nil
}

func weatherToModel(responce *WeatherResponse) *models.CurrentWeather {
	t, _ := responce.CurrentCondition[0].TempC.Float64()
	h, _ := responce.CurrentCondition[0].Humidity.Float64()

	d := models.NoWeatherDescription
	if len(responce.CurrentCondition[0].WeatherDesc) > 0 {
		d = responce.CurrentCondition[0].WeatherDesc[0].Value
	}

	return &models.CurrentWeather{
		Temperature: t,
		Humidity:    h,
		Description: d,
	}
}
