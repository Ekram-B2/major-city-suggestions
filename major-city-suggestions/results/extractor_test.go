package results

import (
	"reflect"
	"testing"
)

func Test_results_GetCitySetFromDataset(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	var expectedOut interface{}
	// 3. Define input arg (Act)
	var data interface{}
	dataSet := map[string]interface{}{"cities": data}
	// 4. Determine output from function (Act)
	actualOut, err := GetCitySetFromDataset(dataSet)

	// 5. Determine if expected and actual match (Assert)
	if err != nil {
		t.Fatalf("expected did not match actual")
	}

	if !reflect.DeepEqual(expectedOut, actualOut) {
		t.Fatalf("expected did not match actual")
	}
}
