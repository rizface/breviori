package httpserver

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
)

type deps struct {
	Shortener
}

type HttpServer struct {
	R *chi.Mux
	deps
}

type RoutesRegistrar func(r *chi.Mux)

func NewHTTPServer(rec Recorder) *HttpServer {
	return &HttpServer{
		R: NewRouter(rec),
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
		slog.Error(err.Error())
	}
}
