package wttr_test

import (
	"context"
	"testing"

	"go-weather/app/internal/clients/wttr"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetWeather_Real(t *testing.T) {
	client := wttr.New()

	ctx := context.Background()
	city := "Kyiv"

	result, err := client.GetWeather(ctx, city)

	require.NoError(t, err, "expected no error from real API")
	require.NotNil(t, result, "result should not be nil")

	assert.InDelta(t, -50.0, result.Temperature, 100.0, "temperature should be within realistic range")
	assert.InDelta(t, 0.0, result.Humidity, 100.0, "humidity should be within 0-100")
	assert.NotEmpty(t, result.Description, "description should not be empty")
}
