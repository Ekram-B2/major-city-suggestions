package dataset

import (
	"reflect"
	"testing"
)

func testManifestGetter() string {
	return "dataset/manifest/manifest.test.go"
}

func testDataSetBuilder(manifest Manifest) map[string][]string {
	return map[string][]string{"json": []string{"files/json/ca.json"}}
}

func Test_dataset_LoadPersistanceFiles(t *testing.T) {
	// 1. Nothing to set up for testing (Arrange)

	// 2. Define expected output (Act)
	expectedDataSet := map[string][]string{"json": []string{"files/json/ca.json"}}

	// 3. Define arugments for the test (Act)
	manifest := Manifest{Files: []string{"files/json/ca.json"}}

	// 4. Perform operation (Act)
	actualDataSet := defaultbuildDataSetFrom(manifest)

	// 5. See if the actual matches with what was expected
	if !reflect.DeepEqual(expectedDataSet, actualDataSet) {
		t.Fatalf("the actualData set does match the expected; the actual is %v and the expected is %v", actualDataSet, expectedDataSet)
	}
}
