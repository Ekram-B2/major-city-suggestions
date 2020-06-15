package results

import "testing"

func Test_results_FindCityLatitude(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	expectedOut := "-4.5432"
	// 3. Peform operation (Act)
	actualOut := FindCityLatitude(city{Lat: "-4.5432"})
	// 4. Determine if output matches whatwas expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("expected did not match actul; expected was %s and actual was %s", expectedOut, actualOut)
	}

}

func Test_results_FindCityLongitude(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	expectedOut := "-4.5432"
	// 3. Peform operation (Act)
	actualOut := FindCityLongitude(city{Long: "-4.5432"})
	// 4. Determine if output matches whatwas expected (Assert)
	if actualOut != expectedOut {
		t.Fatalf("expected did not match actul; expected was %s and actual was %s", expectedOut, actualOut)
	}
}

func Test_results_FindLongitude(t *testing.T) {
	tests := []struct {
		name       string
		d1         DataPoint
		longFinder LongFinder
		want       string
	}{
		// 1. Set up what is necessary to run test (Arrange)
		{
			d1:         city{Long: "123"},
			name:       "FindLongitudeCity",
			want:       "123",
			longFinder: FindCityLongitude,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute the output and determine if output matches well with expected (Act & Assert)
			if got := tt.longFinder(tt.d1); got != tt.want {
				t.Errorf("GetDataPointType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_results_FindLatitude(t *testing.T) {
	tests := []struct {
		name      string
		d1        DataPoint
		latFinder LatFinder
		want      string
	}{
		// 1. Set up what is necessary to run test (Arrange)
		{
			d1:        city{Lat: "123"},
			name:      "FindLatitudeCity",
			want:      "123",
			latFinder: FindCityLatitude,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute the output and determine if output matches well with expected (Act & Assert)
			if got := tt.latFinder(tt.d1); got != tt.want {
				t.Errorf("GetDataPointType() = %v, want %v", got, tt.want)
			}
		})
	}
}
