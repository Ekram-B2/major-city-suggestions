package results

import "testing"

func Test_results_ConvertSampleToDataPoint(t *testing.T) {
	tests := []struct {
		name string

		expectedDataPoint         DataPoint
		expectedMinimalProperties []string

		sample    interface{}
		dataPoint string

		dataPointConverter dataPointConverter
	}{
		// 1. Create setup required to run test (Arrange)
		{
			name: "ConvertSampleToDataPointCity",

			expectedDataPoint:         city{Lat: "123", Long: "4556"},
			expectedMinimalProperties: []string{"lat", "long"},
			sample:                    map[string]interface{}{"lat": "123", "long": "4556"},
			dataPoint:                 "city",

			dataPointConverter: ConvertCitySampleToDataPoint,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actualMinimalProperties, actualDataPoint := tt.dataPointConverter(tt.sample, tt.dataPoint)
			if len(actualMinimalProperties) != len(tt.expectedMinimalProperties) {
				t.Fatalf("was not able to calculate the expected minimal properties; actual Minimal Properties is %s and expectedMinimalProperties is %s", actualMinimalProperties, tt.expectedMinimalProperties)
			}

			if !actualDataPoint.Equals(tt.expectedDataPoint) {
				t.Fatalf("was not able to calculate the expected data point; actual data point is %v and expected data point is %v", actualDataPoint, tt.expectedDataPoint)
			}

		})
	}
}
