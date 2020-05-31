package suggestions

import (
	"github.com/major-city-suggestions/datastore"
	"github.com/major-city-suggestions/score"
)

// Suggestion is the transformed output that is returned to the user
type Suggestion struct {
	city  datastore.LargeCity
	score float32
}

// GetSuggestionsForSearchTerm returns a list of suggestions given the input of large cities
func getSuggestionsForSearchTerm(dataState *datastore.DataState, searchTerm string) []Suggestion {
	// 0. Create container to store the suggestions to be returned
	var suggestions []Suggestion

	// 1. Create a suggestion from a city and add it to the wrapper
	for _, relevantCity := range dataState.Cities {
		newScore, err := score.CalculateRelevancyScore(searchTerm, relevantCity)
		if err != nil {
			// This logic is run when we are unable to calculate a score for a city.
			continue
		}
		newSuggestion := Suggestion{city: relevantCity, score: newScore}
		suggestions = append(suggestions, newSuggestion)
	}
	// 2. perform relevancy sort in order to set up the order within which suggestions will be returned back to the user
	suggestions = relevancySort(suggestions)

	// 3. return list of suggestions now within the order that they will be returned
	return suggestions

}

// algorithm to perform a sort on a linear datastructure of suggestions
func relevancySort(suggestions []Suggestion) []Suggestion {
	// set up a bubble sort for now
	for end := len(suggestions) - 1; end > 0; end-- {
		for index := 0; index < end; index++ {
			if suggestions[index].score > suggestions[index+1].score {
				swap(suggestions, index, index+1)
			}
		}
	}
	return []Suggestion{}
}

// algorithm to swap within a linear datastructure of suggestions
func swap(suggestions []Suggestion, i, j int) {
	tempSuggestion := suggestions[i]
	suggestions[i] = suggestions[j]
	suggestions[j] = tempSuggestion
}
