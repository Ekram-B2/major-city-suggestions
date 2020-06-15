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

func testManifestGetter() string {
	return "dataset/manifest/manifest.test.go"
}

func testDataSetBuilder(manifest Manifest) []string {
	return []string{"files/json/ca.json"}
}

func Test_config_buildDataSet(t *testing.T) {
	tests := []struct {
		name                       string
		want                       []string
		manifest                   Manifest
		buildDataSetRepresentation DataSetBuilder
	}{
		// 1. Initialize the parts of the application required to perform the test (Arrange)
		{
			name:                       "buildDataSetDefaultDatasetBuilder",
			want:                       []string{"files/ca.json"},
			manifest:                   manifestDefault{Files: []string{"files/ca.json"}},
			buildDataSetRepresentation: buildDataSetFilesDefault,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Perform operation (Act)
			actualDataSet := tt.buildDataSetRepresentation(tt.manifest)

			// 3. See if the actual matches with what was expected (Assert)
			if !reflect.DeepEqual(tt.want, actualDataSet) {
				t.Fatalf("the actualData set does match the expected; the actual is %v and the expected is %v", actualDataSet, tt.want)
			}
		})
	}
}
