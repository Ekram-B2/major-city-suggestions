package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"

	"github.com/major-city-suggestions/rankmanager-microservice/rankmanager"
)

func main() {
	// 1. Set up router object to define paths which wrap execution logic
	r := chi.NewRouter()

	// 2. define the endpoints required of for the task
	r.Get("/determineRank", rankmanager.HandleRequestToDetermineRank)
	// r.Get("/detemineRankWithLatLng")

	// 3. start up an http server object at port 8080
	fmt.Println(http.ListenAndServe(":8080", r))
}
