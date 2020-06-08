package results

import "testing"

func Test_results_ConvertSampleToDataPoint(t *testing.T) {
	// 1. Nothing to init to set up test (Arrange)

	// 2. Define expected output (Act)
	expectedMinimalProperties := []string{"lat", "lng"}
	expectedDataPoint := city{Lat: "123", Lng: "4556"}

	// 3. Define input args (Act)
	var sample interface{}
	sample = map[string]interface{}{"lat": "123", "lng": "4556"}
	datapoint := city{Lat: "123", Lng: "4556"}
	minimalProperties := []string{"lat", "lng"}

	// 4. Determine output for function (Act)
	actualMinimalProperties, actualDataPoint := ConvertSampleToDataPoint(sample, datapoint, minimalProperties)

	// 5. Determine if actual matches expected
	if len(actualMinimalProperties) != len(expectedMinimalProperties) {
		t.Fatalf("was not able to calculate the expected minimal properties; actualMinimalProperties is %s and expectedMinimalProperties is %s", actualMinimalProperties, expectedMinimalProperties)
	}

	for _, key := range actualMinimalProperties {
		if !isAMember(key, expectedMinimalProperties) {
			t.Fatalf("was not able to calculate the expected minimal properties; actualMinimalProperties is %s and expectedMinimalProperties is %s", actualMinimalProperties, expectedMinimalProperties)
		}
	}

	if expectedDataPoint.Lat != GetLatitudeForDataPoint("city")(actualDataPoint) || expectedDataPoint.Lng != GetLongitudeForDataPoint("city")(actualDataPoint) {
		t.Fatalf("was not able to calculate the expected datapoint;")
	}
}
