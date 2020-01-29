package app

import (
	"net/http"

	"github.com/braddle/go-http-template/rest"
	"github.com/gorilla/mux"
)

type App struct {
	r *mux.Router
}

func (a *App) init() {
	a.r.Handle("/healthcheck", a.getHealthCheckHandle()).Methods(http.MethodGet)
}

func (a *App) getHealthCheckHandle() http.Handler {
	return rest.HealthCheck{}
}

func New(h *mux.Router) *App {
	a := &App{
		r: h,
	}

	a.init()

	return a
}
