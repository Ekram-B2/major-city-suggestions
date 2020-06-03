package datastore

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	l4g "github.com/alecthomas/log4go"
)

var filePath string = "datastore/data/ca.json"

// JSONFileManager is a manager for persistant storage stored at a file
type JSONFileManager struct{}

// GetAllRelevantCities is the interface used to access all relevant search terms
// The data state is stored in persistance as a json file. The output is returned as
// a pointer to because it is possible for the state to have no relevant entries relative
// to the input.
func (fm *JSONFileManager) GetAllRelevantCities(searchTerm string) (*DataState, error) {

	// 1. Retreive the entire loaded datastate
	dataState, err := fm.getAllCities()
	if err != nil {
		// This is a serious error since this means that we are unable to get
		// to the point where the search term is applied onto the suggestions
		l4g.Error("Unable to load state to process.")
		return nil, err
	}

	// 2. Filter away irrelevant items from the DataState
	dataState = fm.filterForRelevantCities(dataState, searchTerm)

	// 3. return the new datastate
	return dataState, nil

}

// getAllCities is an unexported private function whose purpose is to load in
// the datastate from a file stored the server's secondary storage.
func (fm *JSONFileManager) getAllCities() (*DataState, error) {

	// 0. Create container to store large cities
	var cities DataState

	// 1. Open the file storing the state information that we transform
	dataStateBuffer, err := os.Open(filePath)
	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		l4g.Error(err.Error())
		return nil, err
	}

	defer dataStateBuffer.Close()

	// 2. Extract a byte stream from the dataStateBuffer
	byteStream, err := ioutil.ReadAll(dataStateBuffer)
	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		l4g.Error(err.Error())
		return nil, err
	}

	// 3. Unmarshall the byte stream to fit a go structure representation of the byte steam that we can manipulate
	err = json.Unmarshal(byteStream, &cities)
	if err != nil {
		l4g.Error(err.Error())
		return nil, err
	}

	// 3. Return the cities
	return &cities, nil

}

// GetAllRelevantCitiesLatLng is the interface used to access all relevant search terms
// relative to the provided latitude, longitude. and a search term. The data state is returned as a pointer
// because it is possible for the state to have entries relative to the input.
func (fm *JSONFileManager) GetAllRelevantCitiesLatLng(searchTerm string, latitude, longitude float32) (*DataState, error) {

	// 3. Return the cities
	return &DataState{}, nil
}

// filterForRelevantCities is an unexported private function that filters entries away from
// from the DataState that are irrelevant to the search term
func (fm *JSONFileManager) filterForRelevantCities(dataState *DataState, searchTerm string) *DataState {
	// 1. Create container to store the entries that are determined to be relevant
	var newDataState DataState

	// 2. Apply algorithm on each entry and if deemed relevant, add it to the relevant entry container
	for _, city := range dataState.Cities {
		if fm.isRelevant(searchTerm, city) {
			newDataState.Cities = append(newDataState.Cities, city)
		}
	}

	// 3. Return the modified data state
	return &newDataState

}

// isRelevant is the baseline algorithm used to determine if a city is relevant or not
func (fm *JSONFileManager) isRelevant(searchTerm string, city LargeCity) bool {
	return strings.ContainsAny(searchTerm, city.City)
}
