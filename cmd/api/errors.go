package main

import (
	"fmt"
	"net/http"
)

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Page not found."))
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusMethodNotAllowed)
	w.Write([]byte(fmt.Sprintf("Method \"%s\" is not allowed.", r.Method)))
}
