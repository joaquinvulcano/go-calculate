package main

import (
	"context"
	"flag"
	"log/slog"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/rs/cors"
)

type ctxKey string

const requestIDKey ctxKey = "requestID"

func main() {
	cliMode := flag.Bool("cli", false, "Ejecutar en modo CLI")
	flag.Parse()

	if *cliMode {
		startCLI()
		return
	}

	// Configurar el logger
	logger := slog.Default()

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPost},
		AllowedHeaders: []string{"Content-Type", "Authorization"},
	})

	// Crear el enrutador
	mux := http.NewServeMux()

	// Registrar los handlers
	mux.HandleFunc("/calculate", calculateHandler)

	// Configurar el servidor HTTP
	server := &http.Server{
		Addr:         ":8080",
		Handler:      c.Handler(loggingMiddleware(requestIDMiddleware(mux))),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	logger.Info("Starting server on :8080")
	if err := server.ListenAndServe(); err != nil {
		logger.Error("Server error", "error", err)
	}
}

func requestIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestID := uuid.New().String()
		ctx := context.WithValue(r.Context(), requestIDKey, requestID)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger := slog.Default().With(
			"method", r.Method,
			"path", r.URL.Path,
			"requestID", r.Context().Value(requestIDKey),
		)

		logger.Info("Request started")

		next.ServeHTTP(w, r)

		logger.Info("Request completed",
			"duration", time.Since(start),
		)
	})
}
