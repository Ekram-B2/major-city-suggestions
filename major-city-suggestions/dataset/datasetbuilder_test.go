package dataset

import (
	"reflect"
	"testing"
)

func Test_dataset_getExtension(t *testing.T) {
	// 1. Nothing to init to set up writing the test (Arrange)

	// 2. Set up expected value (Act)
	expectedExt := "json"

	// 3. Define input file (Act)
	path := "/dataset/data.json"

	// 4. Apply operation to generate output (Act)
	actualExt := getExtension(path)

	// 5. Determine if actual matches with what was expected (Assert)
	if actualExt != expectedExt {
		t.Fatalf("the actual ext does not match with the expected ext; expected is %s and actual is %s", expectedExt, actualExt)
	}

}

func Test_dataset_defaultbuildDataSetFrom(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	expectedDataSet := map[string][]string{"json": []string{"ca.json"}}

	// 3. Define arugments for the test (Act)
	manifest := Manifest{Files: []string{"ca.json"}}

	// 4. Perform operation (Act)
	actualDataSet := defaultbuildDataSetFrom(manifest)

	// 5. See if the actual matches with what was expected
	if !reflect.DeepEqual(expectedDataSet, actualDataSet) {
		t.Fatalf("the actualData set does match the expected; the actual is %v and the expected is %v", actualDataSet, expectedDataSet)
	}
}
