package suggestions

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
	"strings"

	l4g "github.com/alecthomas/log4go"

	"github.com/Ekram-B2/suggestionsmanager/config"
	"github.com/Ekram-B2/suggestionsmanager/dataset"
	"github.com/Ekram-B2/suggestionsmanager/relevantreader"
	"github.com/Ekram-B2/suggestionsmanager/results"
)

var (
	noSuggestionsFound int = 0
)

type responseFormat struct {
	Suggestions []Suggestion `json:"suggestions"`
}

func getReader(config config.Config) relevantreader.RelevantReader {
	if !config.IsRemote {
		return relevantreader.NewRelevantFileReader(config,
			dataset.GetDatasetBuilder(config.DataSetBuildType),
			dataset.GetDataSetLoader(config.DataSetLoaderType))
	}
	// This would nominally be the case where a reader would be created to support access to remote clients (e.g. sqldb) but this
	// implementation presently doesn't support that so we return the same result as the local client for now
	return relevantreader.NewRelevantFileReader(config,
		dataset.GetDatasetBuilder(config.DataSetBuildType),
		dataset.GetDataSetLoader(config.DataSetLoaderType))

}

// HandleRequestForSuggestions handles the logic used to build the list of suggestions to return back the the caller
func HandleRequestForSuggestions(rw http.ResponseWriter, req *http.Request) {

	// 1. Check to see the request made containers the required query parameters for performing the search
	searchTerm := req.URL.Query().Get("q")
	if searchTerm == "" {
		l4g.Error("SYSTEM-ERROR: unable to find required query parameter 'q'")
		http.Error(rw, "There was an error retrieving one of the required query parameters; please include a parameter for 'q' in your request.", http.StatusBadRequest)
		return
	}

	// 2. Check to see if the required query parameters are provided
	searchTermLat := req.URL.Query().Get("latitude")

	searchTermLng := req.URL.Query().Get("longitude")

	// 3. Parse to make sure lat and long input are properly formatted and are latitudes and longitudes
	if strings.Contains(searchTermLat, "%20") {
		l4g.Error("SYSTEM-ERROR: invalid parameter 'latitude' provided")
		http.Error(rw, "There was an error processing one of the passed parameters. If you will provide a 'latitude', please provide one in the range [-90,90].", http.StatusBadRequest)
		return
	}

	if strings.Contains(searchTermLng, "%20") {
		l4g.Error("SYSTEM-ERROR: invalid parameter 'longitude' provided")
		http.Error(rw, "There was an error retreiving one of the required query parameters. If you will provide a 'longitude', please provide one in the range [-180,180]", http.StatusBadRequest)
		return
	}

	isFormattedWell := verifyLatLongIsWellFormatted(searchTermLat, searchTermLng, rw)
	if isFormattedWell != true {
		return
	}
	// 4. Load in system configuration for project
	config, err := config.LoadConfiguration(config.GetConfigPath(os.Getenv("CONFIG_OPERATION_TYPE")))
	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to load in configuration object: %s", err.Error())
		http.Error(rw, "There was an issue on our end; please wait for some time before trying your request again.", http.StatusInternalServerError)
		return
	}
	// 5. Create a reader to apply for reading in structured data
	reader := getReader(config)

	// 6. Read in a set of structured results from the data source
	structuredResults, err := reader.ReadRelevant(searchTerm)

	if err != nil {
		l4g.Error("OPERATION-ERROR: error in process applied to convert persistance to structured output: %s", err.Error())
		http.Error(rw, "There was an issue on our end; please wait for some time before trying your request again.", http.StatusInternalServerError)
		return
	}

	// 7. Transform the structured result set into a list of suggestions - i.e, the form that we want to return them
	suggestions := convertResultsIntoSugestions(structuredResults, searchTerm,
		results.GetLatitudeForDataPoint(config.DataPointType),
		results.GetLongitudeForDataPoint(config.DataPointType),
		searchTermLat,
		searchTermLng)

	// 8. Perform a sort upon the suggestions
	applyRelevancySorter(config.SorterType)(suggestions)

	// 9. Set up a formatted response to be returned back to the caller
	responseContent := responseFormat{Suggestions: suggestions}

	// 10. Return response back to the caller
	rw.Header().Add("Content-Type", "application/json; charset=UTF-8")
	rw.WriteHeader(http.StatusOK)

	b := &bytes.Buffer{}
	err = json.NewEncoder(b).Encode(responseContent)
	if err != nil {
		l4g.Error("OPERATION-ERROR: error in marshalling to response to a output stream: %s", err.Error())
		http.Error(rw, "There was an error marshalling to the expected output; please try again later after waiting some time.", http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(b.Bytes())
	if err != nil {
		l4g.Error("OPERATION-ERROR: error in writing steam out to reponse content: %s", err.Error())
		http.Error(rw, "There was an writing to replying content to the response; please try again later after waiting some time.", http.StatusInternalServerError)
		return
	}
}

// verifyLatLongIsWellFormatted is applied to determine if the lat and lng paramters are properly formatted
func verifyLatLongIsWellFormatted(lat, lng string, rw http.ResponseWriter) (isFormattedWell bool) {
	if lat == "" && lng == "" {
		// in this case nothing was provided
		return true
	}
	if strings.Contains(lat, "%20") {
		l4g.Error("SYSTEM-ERROR: invalid parameter 'latitude' provided")
		http.Error(rw, "There was an error processing one of the passed parameters. If you will provide a 'latitude', please provide one in the range [-90,90].", http.StatusBadRequest)
		return false
	}

	if strings.Contains(lng, "%20") {
		l4g.Error("SYSTEM-ERROR: invalid parameter 'longitude' provided")
		http.Error(rw, "There was an error retreiving one of the required query parameters. If you will provide a 'longitude', please provide one in the range [-180,180]", http.StatusBadRequest)
		return false
	}

	lat64, err := strconv.ParseFloat(lat, 64)
	if err != nil {
		l4g.Error("SYSTEM-ERROR: error when trying to convert the latitude to a float")
		http.Error(rw, "There was an error processing one of the passed parameters. If you will provide a 'latitude', please provide one in the range [-90,90].", http.StatusBadRequest)
		return false
	}

	if lat64 < -90.0 || lat64 > 90.0 {
		l4g.Error("SYSTEM-ERROR: error when checking bounds of the latitude")
		http.Error(rw, "There was an error processing one of the passed parameters. If you will provide a 'latitude', please provide one in the range [-90,90].", http.StatusBadRequest)
		return false
	}

	lng64, err := strconv.ParseFloat(lng, 64)
	if err != nil {
		l4g.Error("SYSTEM-ERROR: error when trying to convert longitude to a float")
		http.Error(rw, "There was an error retreiving one of the required query parameters. If you will provide a 'longitude', please provide one in the range [-180,180]", http.StatusBadRequest)
		return false
	}

	if lng64 < -180.0 || lng64 > 180.0 {
		l4g.Error("SYSTEM-ERROR: error when checking bounds of the longitude")
		http.Error(rw, "There was an error retreiving one of the required query parameters. If you will provide a 'longitude', please provide one in the range [-180,180]", http.StatusBadRequest)
		return false
	}
	return true
}
