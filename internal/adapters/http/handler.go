	package http

import (
    "encoding/json"
    "net/http"
    "strconv"
    "github.com/gastonarias/weather-api/internal/application"
)

type Handler struct {
    service *application.WeatherService
}

func NewHandler(s *application.WeatherService) *Handler {
    return &Handler{service: s}
}

func (h *Handler) Health(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("ok"))
}

func (h *Handler) GetWeather(w http.ResponseWriter, r *http.Request) {
    lat, _ := strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
    lon, _ := strconv.ParseFloat(r.URL.Query().Get("lon"), 64)

    result, err := h.service.GetWeather(r.Context(), lat, lon)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    json.NewEncoder(w).Encode(result)
}