package results

import (
	"testing"
)

type testDataPoint struct{}

func (t testDataPoint) GetDataPointType() string {
	return ""
}

func (t testDataPoint) CanBeCreatedFrom([]string) bool {
	return false
}

func (t testDataPoint) GetStateMutators() map[string]mutator {
	return nil
}

func (t testDataPoint) Equals(DataPoint) bool {
	return false
}

func Test_city_CanBeCreatedFrom(t *testing.T) {
	type fields struct {
		City             string
		Admin            string
		Country          string
		PopulationProper string
		ISO2             string
		Capital          string
		Lat              string
		Lng              string
		Population       string
		containsData     bool
	}
	type args struct {
		foundProperties []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name: "Success",
			args: args{foundProperties: []string{"city", "country", "is02", "lat", "lng"}},
			want: true,
		},
		{
			name: "Failure",
			args: args{foundProperties: []string{"city", "country", "is02", "lat"}},
			want: false,
		},
		{
			name: "Mismatch",
			args: args{foundProperties: []string{"city", "country", "is02", "lng", "lat"}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a city object (Arrange)
			c := city{
				City:             tt.fields.City,
				Admin:            tt.fields.Admin,
				Country:          tt.fields.Country,
				PopulationProper: tt.fields.PopulationProper,
				ISO2:             tt.fields.ISO2,
				Capital:          tt.fields.Capital,
				Lat:              tt.fields.Lat,
				Lng:              tt.fields.Lng,
				Population:       tt.fields.Population,
				containsData:     tt.fields.containsData,
			}
			// Compute the output and determine if output matches well with expected (Act & Assert)
			if got := c.CanBeCreatedFrom(tt.args.foundProperties); got != tt.want {
				t.Errorf("city.CanBeCreatedFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_results_GetSource(t *testing.T) {
	// Create a city object (Arrange)
	city := city{}

	// Define expected output (Act)
	expectedOut := "city"

	// Compute operation for output (Act)
	actualOut := city.GetDataPointType()

	// Determine if the output matches the expected (Assert)
	if expectedOut != actualOut {
		t.Fatalf("did not the expected data point type")
	}
}

func Test_city_Equals(t *testing.T) {

	type fields struct {
		City             string
		Admin            string
		Country          string
		PopulationProper string
		ISO2             string
		Capital          string
		Lat              string
		Lng              string
		Population       string
		containsData     bool
	}
	type args struct {
		d DataPoint
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			args:   args{d: city{City: "toronto"}},
			fields: fields{City: "toronto"},
			name:   "Success",
			want:   true,
		},
		{
			args:   args{d: city{City: "toronto"}},
			fields: fields{City: "montreal"},
			name:   "FailureByContentMisMatch",
			want:   false,
		},
		{
			args:   args{d: testDataPoint{}},
			fields: fields{City: "toronto"},
			name:   "FailureByTypeMisMatch",
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a city object (Arrange)
			c := city{
				City:             tt.fields.City,
				Admin:            tt.fields.Admin,
				Country:          tt.fields.Country,
				PopulationProper: tt.fields.PopulationProper,
				ISO2:             tt.fields.ISO2,
				Capital:          tt.fields.Capital,
				Lat:              tt.fields.Lat,
				Lng:              tt.fields.Lng,
				Population:       tt.fields.Population,
				containsData:     tt.fields.containsData,
			}
			// Compute the output and determine if output matches well with expected (Act & Assert)
			if got := c.Equals(tt.args.d); got != tt.want {
				t.Errorf("city.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}
