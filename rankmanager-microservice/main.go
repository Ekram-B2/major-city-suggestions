package main

import (
	"net/http"

	"github.com/go-chi/chi"

	"github.com/major-city-suggestions/rankmanager-microservice/rankmanager"
)

func main() {
	// 1. Set up router object to define paths which wrap execution logic
	r := chi.NewRouter()

	// 2. define the endpoint required of for the task
	r.Get("/determineRank", rankmanager.HandleRequestToDetermineRank)

	// 3. start up an http server object at port 8080
	http.ListenAndServe(":8080", r)
}
