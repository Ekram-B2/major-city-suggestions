package relevantreader

import (
	"strings"
	"testing"

	"github.com/Ekram-B2/suggestionsmanager/config"
	"github.com/Ekram-B2/suggestionsmanager/dataset"
	"github.com/Ekram-B2/suggestionsmanager/results"
)

func testDataLoader(dataset.DataSetBuilder) ([]string, error) {
	return []string{}, nil
}

func testDataSetBuilder(dataset.Manifest) []string {
	return []string{}
}

func testRelevanceAlgorithm(searchTerm string, dp results.DataPoint) bool {
	return strings.ContainsAny(searchTerm, dp.GetHash())
}

func Test_relevantfilereader_filterForRelevantDataPoints(t *testing.T) {
	// Create a relevantfilereader object (Arrange)
	// in this case, we will set the file type and city, and build a relevancy
	// algorithm accordingly
	config := config.Config{DataPointType: "city"}

	rr := relevantFileReader{config: config}
	// Define valid arguments (Arrange)
	relevanceAlgorithm := testRelevanceAlgorithm
	searchTerm := "tor"
	resultSet := results.Cities{}

	relevantCity := results.GetDataPoint("city")
	irrelevantCity := results.GetDataPoint("city")
	relevantCityMutators := relevantCity.GetStateMutators()
	irrelevantCityMutators := irrelevantCity.GetStateMutators()
	relevantCity = relevantCityMutators["name"]("Toronto")
	irrelevantCity = irrelevantCityMutators["name"]("xyv")

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
