package dataset

import (
	"testing"
)

func isSameSlice(expected, actual []string) bool {
	for index, item := range expected {
		if item != actual[index] {
			return false
		}
	}
	return true
}
func Test_config_GetView(t *testing.T) {
	tests := []struct {
		name     string
		want     []string
		manifest Manifest
	}{
		// 1. Initialize the parts of the application required to perform the test (Arrange)
		{
			name:     "GetViewDefaultManifest",
			want:     []string{"files/ca.json"},
			manifest: manifestDefault{Files: []string{"files/ca.json"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Perform operation to render actual output (Act)
			actual := tt.manifest.GetView()

			// 3. Determine if the actual matches with the expected (Assert)

			if !isSameSlice(actual, tt.want) {
				t.Fatalf("the actual did not match the expected; the actual was %s and the expected was %s", actual, tt.want)
			}
		})
	}
}
