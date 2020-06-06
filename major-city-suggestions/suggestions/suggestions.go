package suggestions

import (
	l4g "github.com/alecthomas/log4go"

	"github.com/major-city-suggestions/major-city-suggestions/rankmanagerclient"
	"github.com/major-city-suggestions/major-city-suggestions/results"
)

// Suggestion is the transformed output presented bacl to the client
type Suggestion struct {
	Name      string  `json:"name"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Score     float32 `json:"score"`
}

// convertResultsIntoSugestions returns a list of suggestions given the input of large cities
func convertResultsIntoSugestions(results results.Results, searchTerm string, latFinder results.LatFinder, lngFinder results.LngFinder) []Suggestion {
	// 0. Create container to store the suggestions to be returned
	var suggestions []Suggestion

	// 1. Create a suggestion from a city and add it to the wrapper
	client := rankmanagerclient.RankManagerClient{}
	for _, dp := range results.GetView() {
		newRank, err := client.GetRank(searchTerm, dp.GetHash())
		if err != nil {
			// This logic is run when we are unable to calculate a score for a city.
			l4g.Error("unable to calculate the rank for this datapoint: %s", err.Error())
			continue
		}
		newSuggestion := Suggestion{Name: dp.GetHash(), Latitude: latFinder(dp), Longitude: lngFinder(dp), Score: newRank.Rank}
		suggestions = append(suggestions, newSuggestion)
	}

	// 2. return suggestions
	return suggestions

}
