package ports

import "context"
import "github.com/gastonarias/weather-api/internal/domain"

type WeatherProvider interface {
    GetWeather(ctx context.Context, lat, lon float64) (domain.Weather, error)
}