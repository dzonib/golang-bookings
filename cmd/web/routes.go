package main

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/dzonib/golang-app-with-templates/pkg/handlers"

	"github.com/go-chi/chi/v5"

	"github.com/dzonib/golang-app-with-templates/pkg/config"
)

func Routes(app *config.AppConfig) http.Handler {
	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(WriteToConsole)
	r.Use(NoSurf)
	r.Use(SessionLoad)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	r.Get("/", handlers.Repo.Home)
	r.Get("/about", handlers.Repo.About)

	return r
}
