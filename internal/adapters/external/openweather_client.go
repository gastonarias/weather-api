package external

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gastonarias/weather-api/internal/domain"
)

type OpenMeteoClient struct{}

func NewOpenMeteoClient() *OpenMeteoClient {
	return &OpenMeteoClient{}
}

func (c *OpenMeteoClient) GetWeather(ctx context.Context, lat, lon float64) (domain.Weather, error) {
	url := fmt.Sprintf(
		"https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true",
		lat, lon,
	)

	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return domain.Weather{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Weather{}, fmt.Errorf("error: %s", resp.Status)
	}

	var data struct {
		CurrentWeather struct {
			Temperature float64 `json:"temperature"`
			Windspeed   float64 `json:"windspeed"`
		} `json:"current_weather"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return domain.Weather{}, err
	}

	return domain.Weather{
		Temperature: data.CurrentWeather.Temperature,
		WindSpeed:   data.CurrentWeather.Windspeed,
		Source:      "open-meteo",
	}, nil
}
