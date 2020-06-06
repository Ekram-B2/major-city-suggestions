package rankmanager

import (
	"bytes"
	"encoding/json"
	"net/http"

	l4g "github.com/alecthomas/log4go"
)

type responseFormat struct {
	Name string  `json:"name"`
	Rank float32 `json:"rank"`
}

// HandleRequestToDetermineRank is the logic used to return a rank of a real term against a search term
func HandleRequestToDetermineRank(rw http.ResponseWriter, req *http.Request) {
	// 1. Check to see if within the request made, there is a query parameter containing the search term
	searchTerm := req.URL.Query().Get("searchTerm")
	if searchTerm == "" {
		l4g.Error("no search term found as a query paramter")
		http.Error(rw, "we were unable to find the required search term for this request; please include a 'searchTerm' parameter in your next request", http.StatusBadRequest)
		return
	}
	// 2. See if the real term is provided
	realTerm := req.URL.Query().Get("realTerm")
	if realTerm == "" {
		l4g.Error("No realTerm was found as a query parameter")
		http.Error(rw, "we were unable to find the required real term parameter for this request; please include a 'realTerm' parameter in your next request", http.StatusBadRequest)
		return
	}

	// 3. Apply the real term and search term recovered to generate a rank
	rank, err := getRankForRealTerm(searchTerm, realTerm, generateRanker("levenstein"))
	if err != nil {
		l4g.Error("unable to get rank for the real term: %s", err.Error())
		http.Error(rw, "we were unable to retreive a the rank for the real term provided; please try again after waiting some time.", http.StatusInternalServerError)
	}

	// 4. Set up reply body to send back to caller
	content := responseFormat{Name: realTerm, Rank: rank}

	// 5. Set up the response object within content to be returned back to the user
	rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)

	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(content)
	if err != nil {
		l4g.Error("there was an error marshalling to the expected output: %s", err.Error())
		http.Error(rw, "we were unable to retreive a rank for the real term provided; please try again after waiting some time.", http.StatusInternalServerError)
		return
	}

	// 6. Write reply content into response object
	_, err = rw.Write(b.Bytes())
	if err != nil {
		l4g.Error("there was an error encoding content within the response writer: %s", err.Error())
		http.Error(rw, "we were unable to return the rank for this city. please try again later after waiting some time.", http.StatusInternalServerError)
		return
	}
}

// // HandleRequestToDetermineRankWithLatLng is the wrapper for all the logic used to get the rank that is
// // to be applied onto the city by considering lattitude and longitude
// func HandleRequestToDetermineRank(rw http.ResponseWriter, req *http.Request) {
// 	// 1. Check to see if within the request made, there is a query parameter containing the search term
// 	searchTerm := req.URL.Query().Get("searchTerm")
// 	if searchTerm == "" {
// 		l4g.Error("No search term detected as a query paramters")
// 		http.Error(rw, "There was an error retreiving the required search term from within the request.", http.StatusBadRequest)
// 		return
// 	}
// 	// 2. See if latitude and longitude are provided as well
// 	latitude := req.URL.Query.Get("lat")
// 	if latitude == "" {

// 	}

// 	http.Error(rw, "There was an error marshalling to the expected output", http.StatusInternalServerError)
// 	return
// }
