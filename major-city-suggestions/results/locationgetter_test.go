package results

import "testing"

func Test_results_FindCityLatitude(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	expectedOut := "-4.5432"
	// 3. Peform operation (Act)
	actualOut := FindCityLatitude(city{Lat: "-4.5432"})
	// 4. Determine if output matches whatwas expected
	if actualOut != expectedOut {
		t.Fatalf("expected did not match actul; expected was %s and actual was %s", expectedOut, actualOut)
	}

}

func Test_results_FindCityLongitude(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	expectedOut := "-4.5432"
	// 3. Peform operation (Act)
	actualOut := FindCityLongitude(city{Lng: "-4.5432"})
	// 4. Determine if output matches whatwas expected
	if actualOut != expectedOut {
		t.Fatalf("expected did not match actul; expected was %s and actual was %s", expectedOut, actualOut)
	}
}
