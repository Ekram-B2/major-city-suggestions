package suggestions

import (
	"testing"
)

func Test_suggestions_swap(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Set up expected output (Act)
	expectedSuggestions := []Suggestion{Suggestion{Name: "1"}, Suggestion{Name: "2"}}

	// 3. Set up input args (Act)
	inputSuggestions := []Suggestion{Suggestion{Name: "2"}, Suggestion{Name: "1"}}

	// 4. Perform the operation (Act)
	swap(inputSuggestions, 0, 1)

	// Check if there if the effect that occured was as what was expected (Assert)
	for i, suggestion := range expectedSuggestions {
		if suggestion.Name != inputSuggestions[i].Name {
			t.Fatalf("expected does not match actual; expected was %v but actual was %v", expectedSuggestions, inputSuggestions)
		}
	}

}

func Test_suggestions_sort(t *testing.T) {
	tests := []struct {
		name                string
		expectedSuggestions []Suggestion
		inputSuggestions    []Suggestion
		sorter              sorter
	}{
		// 1. Set up expected and input suggestions to perform sort and then compare the result of a sort (Arrange)
		{
			name:                "bubbleSort",
			expectedSuggestions: []Suggestion{Suggestion{Name: "1", Score: 2}, Suggestion{Name: "2", Score: 1}},
			inputSuggestions:    []Suggestion{Suggestion{Name: "2", Score: 1}, Suggestion{Name: "1", Score: 2}},
			sorter:              bubbleSort,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Check if there if the effect that occured was as what was expected (Assert)
			bubbleSort(tt.inputSuggestions)
			for i, suggestion := range tt.expectedSuggestions {
				if suggestion.Name != tt.inputSuggestions[i].Name {
					t.Fatalf("expected does not match actual; expected was %v but actual was %v", tt.expectedSuggestions, tt.inputSuggestions)
				}
			}
		})
	}
}
