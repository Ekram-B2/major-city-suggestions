package relevantreader

// Naming conventions is based on those defined by hashicorps consul project
import (
	"strings"
	"testing"

	"github.com/major-city-suggestions/major-city-suggestions/config"
	"github.com/major-city-suggestions/major-city-suggestions/dataset"
	"github.com/major-city-suggestions/major-city-suggestions/results"
)

func testDataLoader() (map[string][]string, error) {
	return map[string][]string{}, nil
}

func testManifestPath() string {
	return "relevantreader/test.json"
}

func testDataSetBuilder(dataset.Manifest) map[string][]string {
	return map[string][]string{}
}
func Test_relevantfilereader_NewRelevantFileReaderCreatedWithSystemConfig(t *testing.T) {
	// 1. There is nothing to init for the arrange step

	// 2. Define valid arguments (Act)
	validFileType := "json"
	validDataset := map[string][]string{}
	validDataPointType := "city"
	config := config.SystemConfig{FileType: validFileType, DataPointType: validDataPointType}
	// 3. Define an expected valid state (Act)

	expectedRR := relevantFileReader{fileType: validFileType, dataPoint: validDataPointType}
	// 4. Compute state (Act)

	actualRR := NewRelevantFileReader(config, testManifestPath, testDataSetBuilder, dataset.LoadPersistanceFiles)
	// 5. Determine if expected state matches actual state (Assert)

	if actualRR == nil {
		t.Fatalf("failed to create a new relevant file reader; returned nil result")
	}

	if actualRR.fileType != expectedRR.fileType {
		t.Fatalf("failed to set file type and create a valid initial state for the file reader")
	}

	if actualRR.dataset == nil {
		t.Fatalf("failed to set dataset and create a valid initial state for the file reader")
	}

}

func Test_relevantfilereader_NewRelevantFileReaderWasNotCreatedWithDefaultConfig(t *testing.T) {
	// 1. There is nothing to init for the arrange step (Arrange)

	// 2. Define valid arguments (Act)
	invalidFileType := "csv"
	validDataPoint := "city"
	// Use an implementation of config
	sc := config.SystemConfig{DataPointType: validDataPoint, FileType: invalidFileType}
	// 3. Compute state (Act)
	actualRR := NewRelevantFileReader(sc, testManifestPath, testDataSetBuilder, dataset.LoadPersistanceFiles)
	// 4. Determine if computed state matches expected (Assert)
	if actualRR != nil {
		t.Fatalf("created a non-nil type when constructor input is not valid; actual is nil")
	}
}

func testRelevanceAlgorithm(searchTerm string, dp results.DataPoint) bool {
	return strings.ContainsAny(searchTerm, dp.GetHash())
}

func Test_relevantfilereader_filterForRelevantDataPoints(t *testing.T) {
	// Create a relevantfilereader object (Arrange)
	// in this case, we will set the file type and city, and build a relevancy
	// algorithm accordingly
	validFileType := "json"
	validDataPoint := "city"
	// use an implementation of config
	sc := config.SystemConfig{DataPointType: validDataPoint, FileType: validFileType}

	// Define valid arguments (Act)
	relevanceAlgorithm := test_relevanceAlgorithm
	searchTerm := "tor"
	resultSet := results.Cities{}

	relevantCity := results.GetDataPoint("city")
	irrelevantCity := results.GetDataPoint("city")
	relevantCityMutators := relevantCity.GetStateMutators()
	irrelevantCityMutators := irrelevantCity.GetStateMutators()
	relevantCity = relevantCityMutators["city"]("Toronto")
	irrelevantCity = irrelevantCityMutators["city"]("xyv")

	resultSet.AddDataPoint(relevantCity)
	resultSet.AddDataPoint(irrelevantCity)

	// Compute output (Act)
	results := rr.filterForRelevantDataPoints(searchTerm, resultSet, relevanceAlgorithm)

	// Determine if the relevance filter algorithm succeeds (Assert)
	var relevantCityExists bool
	var irrelevantCityExists bool

	for _, dp := range results.GetView() {
		if dp.Equals(relevantCity) {
			relevantCityExists = true
		}

		if dp.Equals(irrelevantCity) {
			irrelevantCityExists = true
		}
	}

	if relevantCityExists != true && irrelevantCityExists != false {
		t.Fatalf("unable to sort relevant from irrelvant data points")
	}
}
