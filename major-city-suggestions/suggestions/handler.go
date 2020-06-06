package suggestions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"

	l4g "github.com/alecthomas/log4go"

	"github.com/major-city-suggestions/major-city-suggestions/config"
	"github.com/major-city-suggestions/major-city-suggestions/dataset"
	"github.com/major-city-suggestions/major-city-suggestions/relevantreader"
	"github.com/major-city-suggestions/major-city-suggestions/results"
)

type responseFormat struct {
	Suggestions []Suggestion `json:"suggestions"`
}

func getConfiguration(configType string) (config.Config, error) {
	switch configType {
	case "system":
		config := config.SystemConfig{}
		loadedConfig, err := config.LoadConfiguration()
		if err != nil {
			return config, err
		}
		return loadedConfig, nil
	default:
		config := config.SystemConfig{}
		loadedConfig, err := config.LoadConfiguration()
		if err != nil {
			return config, err
		}
		return loadedConfig, nil
	}
}

// HandleRequestForSuggestions handles the logic used to build the list of suggestions to return back the the caller
func HandleRequestForSuggestions(rw http.ResponseWriter, req *http.Request) {

	// 1. Check to see the request made containers the required query parameters for performing the search
	searchTerm := req.URL.Query().Get("q")
	if searchTerm == "" {
		l4g.Error("unable to find required query parameter 'q'")
		http.Error(rw, "there was an error retreiving one of the required query parameters; plese include a parameter for 'q' in your request :)", http.StatusBadRequest)
		return
	}

	// 2. Load in system configuration for project
	_, err := getConfiguration(os.Getenv("CONFIG"))
	if err != nil {
		l4g.Error("unable to load in configuration object: %s", err.Error())
		http.Error(rw, "there was an issue on our end; please wait for some time before trying your request again :)", http.StatusInternalServerError)
		return
	}

	// 3. Create a reader to apply for reading in structured data
	reader := relevantreader.NewRelevantFileReader("city", "json", dataset.LoadPersistanceFiles)

	structuredResults, err := reader.ReadRelevant(searchTerm)
	if err != nil {
		l4g.Error("error in process applied to convert persistance to structured output: %s", err.Error())
		http.Error(rw, "there was an issue on our end; please wait for some time before trying your request again :)", http.StatusInternalServerError)
		return
	}

	// 4. Given the relevant cities, transform them into suggestions - i.e, the form that we want to return them
	suggestions := convertResultsIntoSugestions(structuredResults, searchTerm, results.FindCityLatitude, results.FindCityLongitude)

	// 5. Set up the response format to be returned back to the caller
	responseContent := responseFormat{Suggestions: suggestions}

	// 6. Set up the response object within content to be returned back to the user
	rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)
	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(responseContent)
	if err != nil {
		l4g.Error("error in marshalling to response to a output stream: %s", err.Error())
		http.Error(rw, "there was an error marshalling to the expected output. please try again later after waiting some time :)", http.StatusInternalServerError)
		return
	}
	// 7. Write reply content into response object
	_, err = rw.Write(b.Bytes())
	if err != nil {
		l4g.Error("error in writing steam out to reponse content: %s", err.Error())
		http.Error(rw, "There was an writing to replying content to the response", http.StatusInternalServerError)
		return
	}
}
