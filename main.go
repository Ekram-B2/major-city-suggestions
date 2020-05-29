package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/major-city-suggestions/suggestions"
)

func main() {
	// 1. Set up router object to define paths which wrap execution logic
	r := chi.NewRouter()
	// main endpoint required of for the task
	r.Get("/suggestions", suggestions.HandleRequestForSuggestions)

	// any extra endpoints will be set up here if there is time

	// set up an http server object at port 8080
	http.ListenAndServe(":8080", r)
}
