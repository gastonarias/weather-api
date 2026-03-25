package main

import (
    "log"
    "net/http"
    "github.com/gastonarias/weather-api/internal/adapters/external"
    httpAdapter "github.com/gastonarias/weather-api/internal/adapters/http"
    "github.com/gastonarias/weather-api/internal/application"
)

func main() {
    provider := external.NewOpenMeteoClient()
    service := application.NewWeatherService(provider)
    handler := httpAdapter.NewHandler(service)

    http.HandleFunc("/health", handler.Health)
    http.HandleFunc("/weather", handler.GetWeather)

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}