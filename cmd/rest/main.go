package main

import (
	"log/slog"
	"os"

	"github.com/rizface/breviori/database"
	"github.com/rizface/breviori/httpserver"
	"github.com/rizface/breviori/observer"
)

func main() {
	metricsRecorder := httpserver.Recorder{
		HttpRequestCounter:   observer.RequestCounter(),
		ResponseTimeObserver: observer.TimeResponseHistogram(),
	}

	if err := observer.RegisterMetrics(
		metricsRecorder.HttpRequestCounter,
		metricsRecorder.ResponseTimeObserver,
	); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	database.StartPG()

	database.Migrate()

	database.StartRedis()

	server := httpserver.NewHTTPServer(metricsRecorder)

	server.RegisterRoutes(httpserver.RegisterRoutes)

	server.Start()
}
