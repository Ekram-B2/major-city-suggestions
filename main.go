package main

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"

	"github.com/Ekram-B2/suggestionsmanager/suggestions"
)

const instructionsForApplication string = `Hi!
Welcome to the Suggestion Manager. I'm here to help you out by delivering suggestions that 
can help you complete what you're looking to find. At this time, I'm smart enough to help you 
suggest possible cities you may have been looking for. To ask me for cities, please use 
the following template: 
- '/suggestions?q=<<CITYNAMEHERE>>'.

You may also optionally provide latitude and longitude parameters by applying the following 
template:
- '/suggestions?q=<<CITYNAMEHERE>>&latitude=<<LATITUDEHERE>>&longitude=<<LONGITUDEHERE>>'.`

const instructionsIfWrongEndpoint string = `Hi!
It looks like you passed in the wrong endpoint. If you instead try to make a request by using
a template like:
- '/suggestions?q=<<CITYNAMEHERE>>' 
and I should be able to help you find some suggestions for some of the cities you're looking for. 
If you know of the latitude or longitude of the city, I might be able to come up with a better list 
of suggestions! If you know them, ask me for cities but using:  
- '/suggestions?q=<<CITYNAMEHERE>>&latitude=<<LATITUDEHERE>>&longitude=<<LONGITUDEHERE>>'.`

func main() {
	// 1. Set up router object to define paths that point to the logic that write to reply
	// headers and data given a request
	r := chi.NewRouter()

	r.Get("/", handleRoot)

	r.Get("/suggestions", suggestions.HandleRequestForSuggestions)

	// 2. Define catch all endpoint to help determine how to recover from the error case
	r.Get("/*", handleCatchAll)

	// 3. Define the binding port and which handler to apply based on the build is a development
	// build or production build

	var bindingPort string
	if os.Getenv("DEPLOYMENT_TYPE") == "1" {
		// Hardcoded the port number in development mode
		bindingPort = ":8080"
	} else {
		bindingPort = ":" + os.Getenv("PORT")

	}
	// 4. Start the web application process and bind the application to a port
	http.ListenAndServe(bindingPort, r)
}

// handleRoot is a handler set up to respond to cases where root endpoint is pinged
func handleRoot(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(instructionsForApplication))
	return
}

// handleCatchAll is a handler set up to response to cases where a user requests from
// an endpoint that cannot be handled by the application as it is
func handleCatchAll(rw http.ResponseWriter, req *http.Request) {
	rw.Write([]byte(instructionsIfWrongEndpoint))
}
