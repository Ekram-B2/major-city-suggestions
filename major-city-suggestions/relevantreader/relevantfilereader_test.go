package relevantreader

// Naming conventions is based on those defined by hashicorps consul project
import (
	"testing"
)

func testDataLoader() (map[string][]string, error) {
	return map[string][]string{}, nil
}
func Test_relevantfilereader_NewRelevantFileReaderCreated(t *testing.T) {
	// 1. There is nothing to init for the arrange step

	// 2. Define valid arguments (Act)

	validFileType := "json"
	validDataset := map[string][]string{}
	validDataPoint := "city"
	// 3. Define an expected valid state (Act)

	expectedRR := relevantFileReader{fileType: validFileType, dataPoint: validDataPoint, dataset: validDataset}
	// 4. Compute state (Act)

	actualRR := NewRelevantFileReader(validDataPoint, validFileType, testDataLoader)
	// 5. Determine if expected state matches actual state (Assert)

	if actualRR == nil {
		t.Fatalf("failed to create a new relevant file reader")
	}

	if actualRR.fileType != expectedRR.fileType {
		t.Fatalf("failed to set file type and create a valid initial state for the file reader")
	}

	if actualRR.dataset == nil {
		t.Fatalf("failed to set dataset and create a valid initial state for the file reader")
	}

}

func Test_relevantfilereader_NewRelevantFileReaderWasNotCreated(t *testing.T) {
	// 1. There is nothing to init for the arrange step (Arrange)

	// 2. Define valid arguments (Act)
	invalidFileType := "csv"
	validDataPoint := "city"
	// 4. Compute state (Act)
	actualRR := NewRelevantFileReader(validDataPoint, invalidFileType, testDataLoader)
	// 5. Determine if computed state matches expected (Assert)
	if actualRR != nil {
		t.Fatalf("created a non-nil type when constructor input is not valid")
	}
}

// func Test_relevantfilereader_isRelevant(t *testing.T) {
// 	// 1. Nothing to init for the arrange step (Arrange)

// 	// 2. Define valid arguments (Act)
// 	searchTerm := "tor"
// 	dp := city{City: "Toronto"}

// 	// 3. Define expected valid output (Act)
// 	expectedRes := true

// 	// 4. Compute output (Act)
// 	actualRes := isRelevant(searchTerm, newdp)

// 	// 5. Determine if computed output matches expected (Assert)
// 	if actualRes != expectedRes {
// 		t.Fatalf("unable to determine if city is relevant")
// 	}

// }

// func Test_relevantfilereader_filterForRelevantDataPoints(t *testing.T) {
// 	// Create a relevantfilereader object (Arrange)
// 	validFileType := "json"
// 	validDataset := map[string][]string{}
// 	validDataPoint := "city"
// 	rr := NewRelevantFileReader(validDataPoint, validFileType, validDataset)

// 	// Define valid arguments (Act)
// 	relevanceAlgorithm := isRelevant
// 	searchTerm := "tor"
// 	resultSet := results.Cities{}

// 	relevantCity := results.GetDataPoint("city")
// 	irrelevantCity := results.GetDataPoint("city")
// 	relevantCityMutators := relevantCity.GetStateMutators()
// 	irrelevantCityMutators := irrelevantCity.GetStateMutators()
// 	relevantCity = relevantCityMutators["city"]("Toronto")
// 	irrelevantCity = irrelevantCityMutators["city"]("xyv")

// 	resultSet.AddDataPoint(relevantCity)
// 	resultSet.AddDataPoint(irrelevantCity)

// 	// Compute output (Act)
// 	results := rr.filterForRelevantDataPoints(searchTerm, resultSet, relevanceAlgorithm)

// 	// Determine if the relevance filter algorithm succeeds (Assert)
// 	var relevantCityExists bool
// 	var irrelevantCityExists bool

// 	for _, dp := range results.GetView() {
// 		if dp.Equals(relevantCity) {
// 			relevantCityExists = true
// 		}

// 		if dp.Equals(irrelevantCity) {
// 			irrelevantCityExists = true
// 		}
// 	}

// 	if relevantCityExists != true && irrelevantCityExists != false {
// 		t.Fatalf("unable to sort relevant from irrelvant data points")
// 	}
// }
