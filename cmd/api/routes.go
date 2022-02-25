package main

import (
	"net/http"

	"github.com/br7552/router"
)

func (app *application) routes() *router.Router {
	router := router.New()
	router.MethodNotAllowed = app.methodNotAllowedResponse
	router.NotFound = app.notFoundResponse

	router.HandleFunc(http.MethodGet, "/", app.addrInfoHandler)
	router.HandleFunc(http.MethodGet, "/ip/:ip", app.meetingHandler)

	return router
}
