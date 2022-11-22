package main

import (
	"net/http"

	"github.com/go-chi/chi"
)

// routes provide the different routes for the application
func (app *application) routes() *chi.Mux {
	router := chi.NewRouter()

	router.NotFound(app.notFoundResponse)
	router.MethodNotAllowed(app.methodNotAllowedResponse)

	router.Get("/healtcheck", app.healtcheckHandler)
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome"))
	})

	router.Get("/actors/{id}", app.showActorHandler)

	return router
}
