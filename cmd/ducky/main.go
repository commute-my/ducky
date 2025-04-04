package main

import (
	"log/slog"
	"net/http"
	"net/url"
	"os"
	"time"

	"github.com/commute-my/ducky/handler"
	"github.com/commute-my/ducky/motis"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/httprate"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	rawMotisUrl := os.Getenv("MOTIS_BASE_URL")
	if rawMotisUrl == "" {
		logger.Error("MOTIS_BASE_URL is not present")
		os.Exit(1)
	}

	motisUrl, err := url.Parse(rawMotisUrl)
	if err != nil {
		logger.Error("could not parse MOTIS_BASE_URL", slog.Any("err", err))
		os.Exit(1)
	}

	motisCli := motis.NewClient(motis.WithBaseUrl(motisUrl))
	geocodeHandler := handler.NewGeocoder(motisCli)
	planHandler := handler.NewPlanner(motisCli)

	rl := httprate.NewRateLimiter(100, time.Minute, httprate.WithKeyByRealIP(), httprate.WithResponseHeaders(httprate.ResponseHeaders{
		Limit:      "X-RateLimit-Limit",
		Remaining:  "X-RateLimit-Remaining",
		Reset:      "X-RateLimit-Reset",
		RetryAfter: "Retry-After",
		Increment:  "",
	}))

	r := chi.NewRouter()
	r.Use(rl.Handler)
	r.Use(httprate.LimitByIP(100, time.Minute))
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
	r.Route("/v1", func(r chi.Router) {
		r.Post("/geocode/search", geocodeHandler.Search)
		r.Post("/plan", planHandler.Plan)
	})

	logger.Info("listening http server on 127.0.0.1:8100")
	srv := &http.Server{
		Addr:    "127.0.0.1:8100",
		Handler: r,
	}
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("could not listen http server", slog.Any("err", err))
		os.Exit(1)
	}
}
