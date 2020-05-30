package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

// LargeCity is the structure that enacts as interface to access APIs
type LargeCity struct {
	City             string `json:"city"`
	Admin            string `json:"admin"`
	Country          string `json:"country"`
	PopulationProper string `json:"population_proper"`
	ISO2             string `json:"iso2"`
	Capital          string `json:"capital"`
	Lat              string `json:"lat"`
	Lng              string `json:"lng"`
	Population       string `json:"population"`
}

// DataState is the structure that represents the underlying state of the system
type DataState struct {
	Cities []LargeCity `json:"cities"`
}

// GetAllRelevantCities is the interface used to access all relevant search terms
// A data state is returned as a pointer to because it is possible for the state to
// have no relevant entries given the search term.
func GetAllRelevantCities(searchTerm string) (*DataState, error) {

	// 1. Retreive loaded datastate
	dataState, err := getAllCities()
	if err != nil {
		// This is a serious error since this means that we are unable to get
		// to the point where the search term is applied onto the suggestions
		return nil, err
	}

	// 2. Filter away irrelevant items from the DataState
	dataState = filterForRelevantCities(dataState, searchTerm)

	// 3. return the new datastate
	return dataState, nil

}

// getAllCities is an unexported private function whose purpose is load in a representation
// of the entire data state in the form of a go structure.
func getAllCities() (*DataState, error) {
	// 0. Create container to store large cities

	var cities DataState

	// 1. Open the file storing the state information that we transform
	dataStateBuffer, err := os.Open("data/ca.json")

	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		return nil, err
	}

	defer dataStateBuffer.Close()

	// 2. Extract a byte stream from the dataStateBuffer
	byteStream, err := ioutil.ReadAll(dataStateBuffer)
	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		return nil, err
	}

	// 3. Unmarshall the byte stream to fit a go structure representation of the byte steam that we can manipulate
	err = json.Unmarshal(byteStream, &cities)
	if err != nil {
		return nil, err
	}

	// 3. Return cities
	return &cities, nil

}

// filterForRelevantCities is an unexported private function that filters entries away from
// from the DataState that are irrelevant to the search term
func filterForRelevantCities(dataState *DataState, searchTerm string) *DataState {
	// 1. Create container to store the entries that are determined to be relevant
	var newDataState DataState
	// 2. Apply algorithm on each entry and if deemed relevant, add it to the relevant entry container
	for _, city := range dataState.Cities {
		if isRelevant(searchTerm, city) {
			newDataState.Cities = append(newDataState.Cities, city)
		}
	}
	// 3. Return the modified data state
	return &newDataState

}

// isRelevant is the baseline algorithm used to determine if a city is relevant or not
func isRelevant(searchTerm string, city LargeCity) bool {
	return strings.ContainsAny(searchTerm, city.City)
}
