package http

import (
	"net/http"

	"github.com/google/uuid"
	"go.uber.org/zap"

	"github.com/gastonarias/weather-api/internal/infrastructure"
)

func RequestIDMiddleware(logger *zap.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		requestID := uuid.New().String()

		ctx := infrastructure.WithRequestID(r.Context(), requestID)
		r = r.WithContext(ctx)

		// opcional: devolverlo al cliente
		w.Header().Set("X-Request-ID", requestID)

		next.ServeHTTP(w, r)
	})
}
