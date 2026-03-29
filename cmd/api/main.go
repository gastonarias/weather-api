package main

import (
	"log"
	"net/http"
	"os"

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

	withLogging := httpAdapter.LoggingMiddleware(logger, mux)
	withRequestID := httpAdapter.RequestIDMiddleware(logger, withLogging)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("starting server", zap.String("port", port))

	if err := http.ListenAndServe(":"+port, withRequestID); err != nil {
		logger.Fatal("server failed", zap.Error(err))
	}
}
