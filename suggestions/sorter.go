package suggestions

type sorter func([]Suggestion)

func applyRelevancySorter(sortType string) sorter {
	switch sortType {
	case "default":
		return bubbleSort
	default:
		return bubbleSort
	}
}

// algorithm to perform a sort on a linear datastructure of suggestions
func bubbleSort(suggestions []Suggestion) {
	// set up a bubble sort for now
	for end := len(suggestions) - 1; end > 0; end-- {
		for index := 0; index < end; index++ {
			if suggestions[index].Score < suggestions[index+1].Score {
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
