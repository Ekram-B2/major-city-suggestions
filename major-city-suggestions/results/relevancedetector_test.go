package results

import "testing"

func Test_relevantfilereader_IsRelevantCity(t *testing.T) {
	// 1. Nothing to init for the arrange step (Arrange)

	// 2. Define valid arguments (Act)
	searchTerm := "tor"
	dp := city{City: "Toronto"}

	// 3. Define expected valid output (Act)
	expectedRes := true

	// 4. Compute output (Act)
	actualRes := IsRelevantCity(searchTerm, dp)

	// 5. Determine if computed output matches expected (Assert)
	if actualRes != expectedRes {
		t.Fatalf("unable to determine if city is relevant; the actual is %v and the expected is %v", actualRes, expectedRes)
	}

}
