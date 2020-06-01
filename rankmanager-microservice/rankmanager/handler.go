package rankmanager

import (
	"bytes"
	"encoding/json"
	"net/http"

	l4g "github.com/alecthomas/log4go"
)

type responseFormat struct {
	CityName string  `json "cityname"`
	Rank     float32 `json "rank"`
}

// HandleRequestToDetermineRank is the wrapper for all the logic used to get the rank that is to
// be applied onto the city
func HandleRequestToDetermineRank(rw http.ResponseWriter, req *http.Request) {
	// 1. Check to see if within the request made, there is a query parameter containing the search term
	searchTerm := req.URL.Query().Get("searchTerm")
	if searchTerm == "" {
		l4g.Error("No search term detected as a query paramters")
		http.Error(rw, "We were unable to find the required search term for this request.", http.StatusBadRequest)
		return
	}
	// 2. See if the city is provided
	city := req.URL.Query().Get("city")
	if city == "" {
		l4g.Error("No city detected as a query parameter")
		http.Error(rw, "We were unable to find the required city parameter from this request.", http.StatusBadRequest)
		return
	}

	// 3. Apply the city and search term recovered to generate a rank
	rank, err := getRankForCity(searchTerm, city, generateRanker(config.ranker))
	if err != nil {
		l4g.Error("Unable to get rank from ranking algrithm: %s", err.Error())
		http.Error(rw, "We were unable to retreive a the rank for this city", http.StatusInternalServerError)
	}

	// 4. Set up reply body to send back to caller
	responseContent := responseFormat{CityName: city, Rank: rank}

	// 5. Set up the response object within content to be returned back to the user
	rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(responseContent)
	if err != nil {
		l4g.Error("There was an error marshalling to the expected output: %s", err.Error())
		http.Error(rw, "We were unable to return the rank for this city", http.StatusInternalServerError)
		return
	}

	// 6. Write reply content into response object
	_, err = rw.Write(b.Bytes())
	if err != nil {
		l4g.Error("There was an error encoding content within the response writer: %s", err.Error())
		http.Error(rw, "We were unable to return the rank for this city", http.StatusInternalServerError)
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
