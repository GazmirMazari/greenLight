package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"greenlight.gazmir.mazari.com/internal/data"

	"github.com/julienschmidt/httprouter"
)

//Add a createMovide handler for the Post method

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Create a new movie")
}

//Add a showMovie handler for the Get method

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {
	//When httprouter is parsing a request it will extract any URL parameters and add them to the request context.

	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	Movie := data.Movie{
		ID:        id,
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"Drama", "Romance", "War"},
		Version:   1,
		CreatedAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, Movie, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
