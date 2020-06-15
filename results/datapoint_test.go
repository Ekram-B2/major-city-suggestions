package results

import (
	"testing"
)

func Test_results_CanBeCreatedFrom(t *testing.T) {

	tests := []struct {
		city            city
		name            string
		foundProperties []string
		want            bool
	}{
		// 1. Init arguments and expectation required to test function (Arrange)
		{
			city:            city{},
			name:            "CanBeCreatedFromCitySuccess",
			foundProperties: []string{"name", "iso2", "lat", "long", "country"},
			want:            true,
		},
		{
			city:            city{},
			name:            "CanBeCreatedFromCityFailure",
			foundProperties: []string{"name", "country", "iso2", "lat"},
			want:            false,
		},
		{
			city:            city{},
			name:            "CanBeCreatedFromCityMismatch",
			foundProperties: []string{"name", "iso2", "country", "long", "lat"},
			want:            true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute the output and determine if output matches well with expected (Act & Assert)
			if got := tt.city.CanBeCreatedFrom(tt.foundProperties); got != tt.want {
				t.Errorf("CanBeCreatedFrom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_results_GetHash(t *testing.T) {
	tests := []struct {
		name string
		d1   DataPoint
		want string
	}{
		// 1. Create two cities objects to compare against (Arrange)
		{
			d1:   city{Name: "Toronto", Country: "CA"},
			name: "GetHashCity",
			want: "Toronto, CA",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute the output and determine if output matches well with expected (Act & Assert)
			if got := tt.d1.GetHash(); got != tt.want {
				t.Errorf("GetHash() = %v, want %v", got, tt.want)
			}
		})
	}
}
func Test_results_GetDataPointType(t *testing.T) {
	tests := []struct {
		name string
		d1   DataPoint
		want string
	}{
		// 1. Create two cities objects to compare against (Arrange)
		{
			d1:   city{},
			name: "GetDataPointCity",
			want: "city",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute the output and determine if output matches well with expected (Act & Assert)
			if got := tt.d1.GetDataPointType(); got != tt.want {
				t.Errorf("GetDataPointType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_results_Equals(t *testing.T) {
	tests := []struct {
		name string
		d1   DataPoint
		d2   DataPoint
		want bool
	}{
		// 1. Create two cities objects to compare against (Arrange)
		{
			d1:   city{Name: "toronto"},
			d2:   city{Name: "toronto"},
			name: "EqualsCity",
			want: true,
		},
		{
			d1:   city{Name: "toronto"},
			d2:   city{Name: "tor"},
			name: "CityAreNotEqual",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Compute the output and determine if output matches well with expected (Act & Assert)
			if got := tt.d1.Equals(tt.d2); got != tt.want {
				t.Errorf("Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_results_GetProperties(t *testing.T) {
	tests := []struct {
		name       string
		d1         DataPoint
		properties []string
	}{
		// 1. Create what's required to set up test(Arrange)
		{
			d1:         city{Name: "toronto"},
			name:       "GetPropertiesCity",
			properties: []string{"name", "lat", "long", "country"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 2. Get output from operation
			props := tt.d1.GetProperties()
			if len(props) != len(tt.properties) {
				t.Fatalf("actual did not match expected; actual was %v and expected was %v", props, tt.properties)
			}
			if !isSameSlice(props, tt.properties) {
				t.Fatalf("actual did not match expected; actual was %v and expected was %v", props, tt.properties)
			}

		})
	}
}

func isSameSlice(one, two []string) bool {
	for index, item := range one {
		if item != two[index] {
			return false
		}
	}
	return true
}
