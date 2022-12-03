package main

import (
	"errors"
	"net/http"

	"github.com/rtransat/go-api-example/internal/data"
)

func (app *application) showActorHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	actor, err := app.models.Actor.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"actor": actor}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createActorHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name   string
		Active data.Bool
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	actor := &data.Actor{
		Name:   input.Name,
		Active: input.Active,
	}

	err = app.models.Actor.Insert(actor)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"actor": actor}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
