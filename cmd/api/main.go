package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

//Declare a string containing the application version number.

const version = "1.0.0"

// Declare a config struct to hold all the configure	options for the application.
type config struct {
	port int
	env  string
}

//Define an application to hold the dependecines for our HTTP handlers, helpers, and middleware.

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config
	//Read the value of the PORT environment variable into the port field of the config struct.
	flag.IntVar(&cfg.port, "port", 4000, "API server port")

	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	// Initialize a new logger which writes messages to the standard out stream, prefixed with the current date and time.
	//prefixed with the current date and time.
	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	//Declare an instance of the application struct, containing the config struct and the logger
	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	//Start the HTTP server.
	logger.Printf("Starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)

}
