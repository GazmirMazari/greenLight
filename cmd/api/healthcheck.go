package main

import (
	"net/http"
)

// Declare a handler which writes a plain-text response with a 200 OK status code.

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	//Create a data map which holds the information that we want to send in the response.
	data := map[string]string{
		"status":     "available",
		"enviroment": app.config.env,
		"version":    version,
	}

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

}
