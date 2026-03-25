package application

import (
    "context"
    "github.com/gastonarias/weather-api/internal/domain"
    "github.com/gastonarias/weather-api/internal/ports"
)

type WeatherService struct {
    provider ports.WeatherProvider
}

func NewWeatherService(p ports.WeatherProvider) *WeatherService {
    return &WeatherService{provider: p}
}

func (s *WeatherService) GetWeather(ctx context.Context, lat, lon float64) (domain.Weather, error) {
    return s.provider.GetWeather(ctx, lat, lon)
}