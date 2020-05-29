package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/major-city-suggestions/suggestions"
)

func main() {
	// 1. Set up router object to define paths which wrap execution logic
	r := chi.NewRouter()

	// this is the main endpoint required of for the task
	r.Get("/suggestions", suggestions.HandleRequestForSuggestions)

	// ... any extra endpoints will be set up here
	if os.Getenv("LOCAL") == "1" {
		r.Get("/", handleRootLocal) // this endpoint is written within for testing purposes
	} else {
		r.Get("/", handleRootProd) // this endpoint is what will be seen within the production environemnt
	}

	// set up logic for case where there will be an empty return

	// set up options for CORS
	corsOptions := cors.Options{
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Accept", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		AllowedOrigins:   []string{"*"},
	}
	c := cors.New(corsOptions)
	r.Use(c.Handler)

	// 3. start up an http server object at port 8080

	http.ListenAndServe(":8080", r)
}

func handleRootLocal(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("The main server is alive an able to catch the ping"))
	return
}

func handleRootProd(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte("Enter /suggestions?q=<<CITYNAMEHERE>> to get started"))
	return
}
