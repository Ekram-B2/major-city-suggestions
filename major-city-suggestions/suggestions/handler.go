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

func getReader(config config.Config) relevantreader.RelevantReader {
	if config.IsRemoteClient() {
		return relevantreader.NewRelevantFileReader(config,
			dataset.GetDatasetBuilderOp(os.Getenv("DEFAULT_DATASET_BUILDER")),
			dataset.LoadPersistanceFiles)
	}
	// This would nominally be the case where a reader would be created to support access to remote clients (e.g. sqldb) but this
	// implementation presently doesn't support that so we return the same result as the remote client for now
	return relevantreader.NewRelevantFileReader(config,
		dataset.GetDatasetBuilderOp(os.Getenv("DEFAULT_DATASET_BUILDER")),
		dataset.LoadPersistanceFiles)

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
	config, err := config.GetConfiguration(os.Getenv("CONFIG"))
	if err != nil {
		l4g.Error("unable to load in configuration object: %s", err.Error())
		http.Error(rw, "there was an issue on our end; please wait for some time before trying your request again :)", http.StatusInternalServerError)
		return
	}

	// 3. Create a reader to apply for reading in structured data
	reader := getReader(config)

	// 4. Read in a set of structured results from the data source
	structuredResults, err := reader.ReadRelevant(searchTerm)
	if err != nil {
		l4g.Error("error in process applied to convert persistance to structured output: %s", err.Error())
		http.Error(rw, "there was an issue on our end; please wait for some time before trying your request again :)", http.StatusInternalServerError)
		return
	}

	// 5. Transform the structured reset set into a list of suggestions - i.e, the form that we want to return them
	suggestions := convertResultsIntoSugestions(structuredResults, searchTerm, results.GetLatitudeForDataPoint(config.GetDataPointType()), results.GetLongitudeForDataPoint(config.GetDataPointType()))

	// 6. Perform a sort upon the suggestions
	applyRelevancySorter("bubble")(suggestions)

	// 7. Set up a formatted response to be returned back to the caller
	responseContent := responseFormat{Suggestions: suggestions}

	// 8. Return formatted response back to the caller
	rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)

	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(responseContent)
	if err != nil {
		l4g.Error("error in marshalling to response to a output stream: %s", err.Error())
		http.Error(rw, "there was an error marshalling to the expected output; please try again later after waiting some time :)", http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(b.Bytes())
	if err != nil {
		l4g.Error("error in writing steam out to reponse content: %s", err.Error())
		http.Error(rw, "there was an writing to replying content to the response; please try again later after waiting some time :)", http.StatusInternalServerError)
		return
	}
}
