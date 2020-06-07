package dataset

import "testing"

func Test_dataset_getManifestPath(t *testing.T) {
	// 1. Nothing to init to set up writing the test (Arrange)

	// 2. Set up expected value (Act)
	expectedOut := "major-city-suggestions/dataset/manifest.json"

	// 3. Apply operation to generate output (Act)
	actualOut := getManifestPath()

	// 5. Determine if actual matches with what was expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("the actual out does not match with the expected out; expected is %s and actual is %s", expectedOut, actualOut)
	}

}
