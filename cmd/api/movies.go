package main

import (
	"fmt"
	"net/http"
	"strconv"

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

	fmt.Fprintln(w, "Show the details of movie %d", id)
}
