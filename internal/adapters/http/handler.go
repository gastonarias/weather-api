package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gastonarias/weather-api/internal/application"
	"github.com/gastonarias/weather-api/internal/infrastructure"
	"go.uber.org/zap"
)

type Handler struct {
	service *application.WeatherService
	logger  *zap.Logger
}

func NewHandler(s *application.WeatherService, logger *zap.Logger) *Handler {
	return &Handler{service: s, logger: logger}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {

	logger := infrastructure.GetLogger(r.Context())

	latStr := r.URL.Query().Get("lat")
	lonStr := r.URL.Query().Get("lon")

	lat, err := strconv.ParseFloat(latStr, 64)
	if err != nil {
		logger.Warn("invalid lat",
			zap.String("lat", latStr),
			zap.Error(err),
		)

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid lat",
		})
		return
	}

	lon, err := strconv.ParseFloat(lonStr, 64)
	if err != nil {
		logger.Warn("invalid lon",
			zap.String("lon", lonStr),
			zap.Error(err),
		)

		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid lon",
		})
		return
	}

	result, err := h.service.GetWeather(r.Context(), lat, lon)
	if err != nil {
		logger.Error("failed to get weather",
			zap.Error(err),
			zap.Float64("lat", lat),
			zap.Float64("lon", lon),
		)

		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(result)
}
