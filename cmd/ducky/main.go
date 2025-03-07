package main

import (
	"log/slog"
	"net/http"
	"net/url"
	"os"

	"github.com/commute-my/ducky/handler"
	"github.com/commute-my/ducky/ors"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	rawOrsUrl := os.Getenv("ORS_BASE_URL")
	if rawOrsUrl == "" {
		logger.Error("ORS_BASE_URL is not present")
		os.Exit(1)
	}

	orsUrl, err := url.Parse(rawOrsUrl)
	if err != nil {
		logger.Error("could not parse ORS_BASE_URL", slog.Any("err", err))
		os.Exit(1)
	}

	orsCli := ors.NewClient(ors.WithBaseUrl(orsUrl))

	mux := http.NewServeMux()
	mux.HandleFunc("POST /directions", handler.NewDirection(orsCli).Direction)

	logger.Info("listening http server on 127.0.0.1:8100")
	srv := &http.Server{
		Addr:    "127.0.0.1:8100",
		Handler: mux,
	}
	if err := srv.ListenAndServe(); err != nil {
		logger.Error("could not listen http server", slog.Any("err", err))
		os.Exit(1)
	}
}
