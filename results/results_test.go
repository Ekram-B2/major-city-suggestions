package results

import (
	"reflect"
	"testing"
)

func Test_results_ContainsMembers(t *testing.T) {
	tests := []struct {
		name string
		r    Results
		want bool
	}{
		// 1. Create results object to run test with (Arrange)
		{
			r:    Cities{containsMembers: true},
			name: "ContainsMembersCities",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Apply operation to generate expected output (Act)
			actual := tt.r.ContainsMembers()

			// 3. Determine if expected matches actual (Assert)
			if tt.want != actual {
				t.Fatalf("expected did not match actual; expected is %v and actual is %v", tt.want, actual)
			}
		})
	}
}

func Test_results_AddDataPoint(t *testing.T) {
	tests := []struct {
		name string
		r    Results
		d1   DataPoint
		want []DataPoint
	}{
		// 1. Create results object to run test with (Arrange)
		{
			r:    Cities{containsMembers: true},
			d1:   city{},
			name: "AddDataPointCities",
			want: []DataPoint{city{}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 2. Apply operation to generate expected output and determine if expected matches actual (Act, Assert)
			tt.r.AddDataPoint(tt.d1)
			for index, dp := range tt.r.GetView() {
				if !reflect.DeepEqual(dp, tt.want[index]) {
					t.Fatalf("expected did not match actual; expected is %v and actual is %v", dp, tt.want)
				}
			}
			if tt.r.ContainsMembers() != true {
				t.Fatalf("ContainsMembers did not return true even though a data point was added")
			}
		})
	}
}

func Test_results_CombineWith(t *testing.T) {
	tests := []struct {
		name string
		r1   Results
		r2   Results
		want []DataPoint
	}{
		// 1. Create two results objects to run test with (Arrange)
		{
			name: "CombineWithCities",
			r1:   Cities{containsMembers: true}.AddDataPoint(city{Lat: "123"}),
			r2:   Cities{containsMembers: true}.AddDataPoint(city{Long: "456"}),
			want: []DataPoint{city{Lat: "123"}, city{Long: "456"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 2. Apply operation to generate expected output and determine if expected matches actual (Act, Assert)
			tt.r1.CombineWith(tt.r2)
			for index, dp := range tt.r1.GetView() {
				if !reflect.DeepEqual(dp, tt.want[index]) {
					t.Fatalf("expected did not match actual; expected is %v and actual is %v", dp, tt.want)
				}
			}
			if tt.r1.ContainsMembers() != true {
				t.Fatalf("ContainsMembers did not return true even though a data point was added")
			}
		})
	}
}

func Test_results_GetView(t *testing.T) {
	tests := []struct {
		name string
		r    Results
		want []DataPoint
	}{
		// 1. Create results object to run test with (Arrange)
		{
			r:    Cities{containsMembers: true}.AddDataPoint(city{Lat: "123"}).AddDataPoint(city{Long: "456"}),
			name: "GetViewCities",
			want: []DataPoint{city{Lat: "123"}, city{Long: "456"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			// 2. Apply operation to generate expected output, and determine if expected matches actual (Act, Arrange)

			for index, dp := range tt.r.GetView() {
				if !reflect.DeepEqual(dp, tt.want[index]) {
					t.Fatalf("expected did not match actual; expected is %v and actual is %v", dp, tt.want)
				}
			}
		})
	}
}
