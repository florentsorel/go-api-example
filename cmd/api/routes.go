package main

import (
	"github.com/go-chi/chi"
)

// routes provide the different routes for the application
func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/healtcheck", app.healtcheckHandler)

	router.Get("/actors/{id}", app.showActorHandler)
	router.Post("/actors", app.createActorHandler)

	return router
}
