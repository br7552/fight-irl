package main

import (
	"net/http"
	"runtime/debug"
)

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request,
	status int, message interface{}) {
	err := app.writeJSON(w, status, envelope{"error": message}, nil)
	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter,
	r *http.Request, err error) {
	app.logError(r, err)
	app.errorResponse(w, r, http.StatusInternalServerError,
		http.StatusText(http.StatusInternalServerError))
}

func (app *application) notFoundResponse(w http.ResponseWriter,
	r *http.Request) {
	app.errorResponse(w, r, http.StatusNotFound,
		http.StatusText(http.StatusNotFound))
}

func (app *application) methodNotAllowedResponse(w http.ResponseWriter,
	r *http.Request) {
	app.errorResponse(w, r, http.StatusMethodNotAllowed,
		http.StatusText(http.StatusMethodNotAllowed))
}

func (app *application) badRequestResponse(w http.ResponseWriter,
	r *http.Request, err error) {
	app.errorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (app *application) logError(r *http.Request, err error) {
	app.logger.Printf(
		"request method: %s, request url: %s\n"+
			"error message: %s\n"+
			"trace: %s\n",
		r.Method, r.URL.String(), err.Error(), string(debug.Stack()))
}
