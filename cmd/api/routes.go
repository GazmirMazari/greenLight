package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	//Initialize new httprouter instance.
	router := httprouter.New()

	//Register the relevant methods and URL patterns for the healthcheck endpoint
	//endpoint.using the handlerfunc() method.

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthCheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/snippets", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/v1/snippets/:id", app.showMovieHandler)

	return router
}
