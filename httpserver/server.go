package httpserver

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type HttpServer struct {
	R *chi.Mux
}

type RoutesRegistrar func(r *chi.Mux)

func NewHTTPServer() *HttpServer {
	return &HttpServer{
		R: chi.NewRouter(),
	}
}

func (h *HttpServer) RegisterRoutes(registrar RoutesRegistrar) {
	registrar(h.R)
}

func (h *HttpServer) Start() {
	port := os.Getenv("BREVIO_PORT")
	if port == "" {
		port = "8000"
	}

	if err := http.ListenAndServe(":"+port, h.R); err != nil {
		slog.Error("httpserver: failed to start server: %v", err)
	}
}
