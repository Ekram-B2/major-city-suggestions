package datastore

// DataState is the structure that represents the underlying state of the system
type DataState struct {
	Cities []LargeCity `json:"cities"`
}

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

// DataManager is an interface to a flexible run time implementation
type DataManager interface {

	// GetAllRelevantCities is the interface used to access all relevant search terms
	// A data state is returned as a pointer to because it is possible for the state to
	// have no relevant entries relative to the search term
	GetAllRelevantCities(string) (*DataState, error)

	// GetAllRelevantCitiesLatLng is the interface used to access all relevant search terms
	// provided that the user had provided the latitude and longitude.
	// The data state is returned as a pointer to because it is possible for the state to
	// have no relevant entries relative to the search term
	GetAllRelevantCitiesLatLng(string, float32, float32) (*DataState, error)
}
