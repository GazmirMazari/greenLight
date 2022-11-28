package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with a 200 OK status code.

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "enviroment: %s", app.config.env)
	fmt.Fprintf(w, "version: %s", version)
}
