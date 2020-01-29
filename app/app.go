package app

import (
	"net/http"
	"time"

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

func (a *App) Run(host string) error {
	srv := http.Server{
		Addr:         host,
		Handler:      a.r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	return srv.ListenAndServe()
}

func New(h *mux.Router) *App {
	a := &App{
		r: h,
	}

	a.init()

	return a
}
