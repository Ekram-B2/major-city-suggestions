package relevantreader

import "github.com/major-city-suggestions/major-city-suggestions/results"

// RelevantReader supports reading relevant data from a persistant store. Relevant data is
// partial segment of the global data set with which a rank can be attributed
type RelevantReader interface {

	// ReadRelevant used to read in relevant data from a persistant store
	ReadRelevant(string) (results.Results, error)
}

type unmarshaller func([]byte, map[string]interface{}) (map[string]interface{}, error)

type relevanceDetector func(string, results.DataPoint) bool
