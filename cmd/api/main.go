package main

import (
	"log"
	"net/http"

	"go.uber.org/zap"

	"github.com/gastonarias/weather-api/internal/adapters/external"
	httpAdapter "github.com/gastonarias/weather-api/internal/adapters/http"
	"github.com/gastonarias/weather-api/internal/application"
	"github.com/gastonarias/weather-api/internal/infrastructure"
)

func main() {

	logger, err := infrastructure.NewLogger()

	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	logger.Info("starting server")

	provider := external.NewOpenMeteoClient()
	service := application.NewWeatherService(provider)
	handler := httpAdapter.NewHandler(service, logger)

	mux := http.NewServeMux()

	mux.HandleFunc("/health", handler.Health)
	mux.HandleFunc("/weather", handler.GetWeather)

	withRequestID := httpAdapter.RequestIDMiddleware(logger, mux)
	withLogging := httpAdapter.LoggingMiddleware(logger, withRequestID)

	logger.Info("starting server", zap.String("port", "8080"))

	if err := http.ListenAndServe(":8080", withLogging); err != nil {
		logger.Fatal("server failed", zap.Error(err))
	}
}
