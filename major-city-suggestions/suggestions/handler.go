package suggestions

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/major-city-suggestions/datastore"
)

type responseFormat struct {
	suggestions []Suggestion
}

// HandleRequestForSuggestions is the wrapper for all the logic used to build
// the list of suggestions to return
func HandleRequestForSuggestions(rw http.ResponseWriter, req *http.Request) {
	// 0. Create list to store suggestions

	// 1. Check to see if within the request made, there is a query parameters containing the search term
	searchTerm := req.URL.Query().Get("q")
	if searchTerm == "" {
		http.Error(rw, "There was an error retreiving the required search term from within the request.", http.StatusBadRequest)
		return
	}

	// 2. Apply the search term in order to render the list of cities that can possibly be suggestions
	dataStateWithRelevantCities, err := datastore.GetAllRelevantCities(searchTerm)
	if err != nil {
		// This is a serious error with how the architecture is meant to be managed and
		// therefore, we can't recover from this
		http.Error(rw, "There was an error determining the suggestions", http.StatusInternalServerError)
		return
	}

	// 3. Given the relevant cities, transform them into suggestions - i.e, the form that we want to return them
	suggestionsForSearchTerm := getSuggestionsForSearchTerm(dataStateWithRelevantCities, searchTerm)

	// 4. Set up the response format to be returned back to the caller
	responseContent := responseFormat{suggestions: suggestionsForSearchTerm}

	// 5. Set up the response object within content to be returned back to the user
	rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(responseContent)
	if err != nil {
		http.Error(rw, "There was an error marshalling to the expected output", http.StatusInternalServerError)
		return
	}

	// 6. Write reply content into response object
	_, err = rw.Write(b.Bytes())
	if err != nil {
		http.Error(rw, "There was an writing to replying content to the response", http.StatusInternalServerError)
		return
	}
}
