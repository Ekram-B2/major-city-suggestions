package dataset

import (
	"reflect"
	"testing"

	"github.com/Ekram-B2/suggestionsmanager/config"
)

func Test_config_loadDataSet(t *testing.T) {
	tests := []struct {
		name           string
		want           []string
		dataSetLoader  DataSetLoader
		dataSetBuilder DataSetBuilder
	}{
		// 1. Initialize the parts of the application required to perform the test (Arrange)
		{
			name:           "loadDataSetDefaultDataSetLoader",
			want:           []string{"files/json/ca.json"},
			dataSetLoader:  LoadDataSetFilesDefault,
			dataSetBuilder: testDataSetBuilder,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 2. Perform operation (Act)
			config := config.Config{ManifestPath: "files/manifest/manifest.test.json"}
			actualDataSet, err := tt.dataSetLoader(config, tt.dataSetBuilder)
			if err != nil {
				t.Fatalf("was not able to load a dataset")
			}

			// 3. See if the actual matches with what was expected (Assert)
			if !reflect.DeepEqual(tt.want, actualDataSet) {
				t.Fatalf("the actualData set does match the expected; the actual is %v and the expected is %v", actualDataSet, tt.want)
			}
		})
	}
}
