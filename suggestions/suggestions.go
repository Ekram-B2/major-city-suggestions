package suggestions

import (
	l4g "github.com/alecthomas/log4go"

	"github.com/Ekram-B2/rankmanager/rank/rankclient"
	"github.com/Ekram-B2/suggestionsmanager/results"
)

// Suggestion is the transformed output presented bacl to the client
type Suggestion struct {
	Name      string  `json:"name"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Score     float32 `json:"score"`
}

// convertResultsIntoSugestions returns a list of suggestions given the input of large cities
func convertResultsIntoSugestions(results results.Results, searchTerm string, latFinder results.LatFinder, longFinder results.LongFinder, searchTermLat, searchTermLong string) []Suggestion {
	// 0. Create container to store the suggestions to be returned
	var suggestions []Suggestion

	// 1. Create a suggestion from a city and add it to the wrapper
	for _, dp := range results.GetView() {
		newRank, err := rankclient.GetRank(searchTerm, dp.GetHash(), latFinder(dp), searchTermLat, longFinder(dp), searchTermLong)
		if err != nil {
			// This logic is run when we are unable to calculate a score for a city.
			l4g.Error("SYSTEM-ERROR: unable to calculate the rank for this datapoint: %s", err.Error())
			continue
		}

		newSuggestion := Suggestion{Name: dp.GetHash(), Latitude: latFinder(dp), Longitude: longFinder(dp), Score: newRank.Rank}
		suggestions = append(suggestions, newSuggestion)
	}

	// 2. return suggestions
	return suggestions

}
