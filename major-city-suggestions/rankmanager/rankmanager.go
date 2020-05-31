package rankmanager

import "github.com/major-city-suggestions/datastore"

// RankManager is the abstract interface used to make design more flexible
type RankManager interface {
	// CalculateRelevancyScore is the algorithm used to calculate a relevancy score
	// without considerations made for latitude and longitude
	CalculateRelevancyScore(string, datastore.LargeCity) (float32, error)

	// CalculateRelevancyScore is the algorithm used to calculate a score relevancy score
	// with considerations made for  latitude and longitude
	CalculateRelevancyScoreWithLatLng(string, float32, float32, datastore.LargeCity) (float32, error)
}

// min is a function to calculate min with int types
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
