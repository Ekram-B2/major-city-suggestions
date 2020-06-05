package suggestions

import (
	l4g "github.com/alecthomas/log4go"

	"github.com/major-city-suggestions/major-city-suggestions/rankmanagerclient"
	"github.com/major-city-suggestions/major-city-suggestions/relevantreader"
)

// Suggestion is the transformed output that is returned to the user
type Suggestion struct {
	Name      string  `json:"name"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Score     float32 `json:"score"`
}

// GetSuggestionsForSearchTerm returns a list of suggestions given the input of large cities
func getSuggestionsForSearchTerm(dataState *relevantreader.Results, searchTerm string) []Suggestion {
	// 0. Create container to store the suggestions to be returned
	var suggestions []Suggestion

	// 1. Create a suggestion from a city and add it to the wrapper
	client := rankmanagerclient.RankManagerClient{}
	for _, relevantCity := range dataState.GetView() {
		newRank, err := client.GetRank(searchTerm, relevantCity.City)
		if err != nil {
			// This logic is run when we are unable to calculate a score for a city.
			l4g.Error("Unable to calculate the rank for this city.")
			continue
		}
		newSuggestion := Suggestion{Name: getSuggestionName(relevantCity), Latitude: relevantCity.Lat, Longitude: relevantCity.Lng, Score: newRank.Rank}
		suggestions = append(suggestions, newSuggestion)
	}

	// 2. perform relevancy sort in order to set up the order within which suggestions will be returned back to the user
	relevancySort(suggestions)
	// 3. return list of suggestions now within the order that they will be returned
	return suggestions

}

// algorithm to perform a sort on a linear datastructure of suggestions
func relevancySort(suggestions []Suggestion) {
	// set up a bubble sort for now
	for end := len(suggestions) - 1; end > 0; end-- {
		for index := 0; index < end; index++ {
			if suggestions[index].Score > suggestions[index+1].Score {
				swap(suggestions, index, index+1)
			}
		}
	}
}

// algorithm to swap within a linear datastructure of suggestions
func swap(suggestions []Suggestion, i, j int) {
	tempSuggestion := suggestions[i]
	suggestions[i] = suggestions[j]
	suggestions[j] = tempSuggestion
}

func getSuggestionName(city datastore.LargeCity) string {
	return city.City + ", " + city.Admin + ", " + city.ISO2
}
