package results

import (
	"reflect"
	"testing"
)

func Test_results_GetCitySetFromDataset(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	var expectedData interface{}
	expectedOut := []interface{}{expectedData}

	// 3. Define input arg (Act)
	dataSet := map[string]interface{}{"cities": expectedData}
	// 4. Determine output from function (Act)
	actualOut, err := GetCitySetFromDataset(dataSet)

	// 5. Determine if expected and actual match (Assert)
	if err != nil {
		t.Fatalf("expected did not match actual; expected was %v and actual was %v", expectedData, actualOut)
	}

	if !reflect.DeepEqual(expectedOut, actualOut) {
		t.Fatalf("expected did not match actual; expected was %v and actual was %v", expectedData, actualOut)
	}
}
