package datastore

import "os"

// LargeCity is the structure that enacts as interface to access APIs
type LargeCity struct {
	name      string
	latitude  float32
	longitude float32
}

// GetAllRelevantCities is the interface used to access all the search terms stored within the flat file
func GetAllRelevantCities(searchTerm string) []LargeCity {
	return []LargeCity{}
}

func populateAllCities(searchTerm string) ([]LargeCity, error) {
	// 0. Create wrapper to store large cities
	var largeCities []LargeCity

	// 1. Open the file storing the state information that we transform
	newFile, err := os.Open("github.com/major-city-suggestions/datastore/data/canadian_cities")
	if err != nil {
		return []LargeCity{}, err
	}
	// 2. Perform a manipulation operation to check if an entry in the file is relevant

	// If the entry is relevant, then we add it to the wrapper storing the large cities
	return largeCities, nil

}
func isRelevant(searchTerm string) bool {
	return false
}

// Any additional APIs can be defined below
