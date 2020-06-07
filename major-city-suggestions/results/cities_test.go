package results

import (
	"reflect"
	"testing"
)

func Testing_results_ContainsMembers(t *testing.T) {
	// 1. Init Cities from which to run test (Arrange)
	cities := Cities{containsMembers: true}

	// 2. Defined expected output (Act)
	expectedOut := true

	// 3. Apply operation to generate expected output (Act)
	actualOut := cities.ContainsMembers()

	// 4. Determine if expected matches actual (Assert)
	if expectedOut != actualOut {
		t.Fatalf("expected did not match actual; expected is %v and actual is %v", expectedOut, actualOut)
	}
}

func Testing_results_AddDataPoint(t *testing.T) {
	// 1. Init Cities from which to run test (Arrange)
	cities := Cities{containsMembers: true}

	// 2. Defined expected output (Act)
	expectedStatusOfContainsMembers := true
	expectedDataPoints := []DataPoint{city{}}

	// 3. Apply operation to generate expected output (Act)
	cities.AddDataPoint(city{})

	// 4. Determine if expected matches actual (Assert)
	for index, dp := range cities.GetView() {
		if !reflect.DeepEqual(dp, expectedDataPoints[index]) {
			t.Fatalf("expected did not match actual; expected is %v and actual is %v", cities.GetView(), expectedDataPoints)
		}
	}

	if cities.ContainsMembers() != expectedStatusOfContainsMembers {
		t.Fatalf("expected did not match actual; expected is %v and actual is %v", expectedStatusOfContainsMembers, cities.GetView())
	}
}

func Testing_results_GetView(t *testing.T) {
	// 1. Init Cities from which to run test (Arrange)
	cities := Cities{containsMembers: true}

	// 2. Defined expected output (Act)
	expectedDataPoints := []DataPoint{city{}}

	// 3. Apply operation to generate expected output (Act)
	view := cities.GetView()

	// 4. Determine if expected matches actual (Assert)
	for index, dp := range cities.GetView() {
		if !reflect.DeepEqual(dp, view[index]) {
			t.Fatalf("expected did not match actual; expected is %v and actual is %v", cities.GetView(), expectedDataPoints)
		}
	}
}

func Testing_results_CombineWith(t *testing.T) {
	// 1. Init Cities from which to run test (Arrange)
	citiesOne := Cities{containsMembers: true}
	citiesTwo := Cities{containsMembers: true}

	// 2. Defined expected output (Act)
	expectedView := []DataPoint{city{Lat: "123"}, city{Lng: "456"}}

	// 3. Apply operation to generate expected output (Act)
	citiesOne.CombineWith(citiesTwo)

	// 4. Determine if expected matches actual (Assert)
	for index, dp := range citiesOne.GetView() {
		if !reflect.DeepEqual(dp, expectedView[index]) {
			t.Fatalf("actual does not match expected")
		}
	}
}
