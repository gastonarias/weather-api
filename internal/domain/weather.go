package domain

type Weather struct {
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
	WindSpeed   float64 `json:"wind_speed"`
	Source      string  `json:"source"`
}
