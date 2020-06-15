package results

import "testing"

func Test_relevantfilereader_IsRelevantCity(t *testing.T) {
	// 1. Nothing to init for the arrange step (Arrange)

	// 2. Define valid arguments (Act)
	searchTerm := "tor"
	dp := city{Name: "Toronto"}

	// 3. Define expected valid output (Act)
	expectedRes := true

	// 4. Compute output (Act)
	actualRes := IsRelevantCity(searchTerm, dp)

	// 5. Determine if computed output matches expected (Assert)
	if actualRes != expectedRes {
		t.Fatalf("unable to determine if city is relevant; the actual is %v and the expected is %v", actualRes, expectedRes)
	}

}

func Test_results_IsRelevant(t *testing.T) {
	tests := []struct {
		name              string
		d1                DataPoint
		searchTerm        string
		relevanceDetector RelevanceDetector
		want              bool
	}{
		// 1. Set up what is necessary to run test (Arrange)
		{
			d1:                city{Name: "toronto"},
			searchTerm:        "toronto",
			name:              "FindLatitudeCitySuccess",
			want:              true,
			relevanceDetector: IsRelevantCity,
		},
		{
			d1:                city{Name: "toronto"},
			searchTerm:        "xyz",
			name:              "FindLatitudeCityFailure",
			want:              false,
			relevanceDetector: IsRelevantCity,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute the output and determine if output matches well with expected (Act & Assert)
			if got := tt.relevanceDetector(tt.searchTerm, tt.d1); got != tt.want {
				t.Errorf("unable to determine if city is relevant; the actual is %v and the expected is %v", got, tt.want)
			}
		})
	}
}
