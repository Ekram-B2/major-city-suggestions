package datastore

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

// Any additional APIs can be defined below
