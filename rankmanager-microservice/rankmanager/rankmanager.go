package rankmanager

import l4g "github.com/alecthomas/log4go"

// rankManager is the abstract interface used to make design more flexible
type rankManager interface {
	// CalculateRelevancyScore is the algorithm used to calculate a relevancy score
	// without considerations made for latitude and longitude
	calculateRelevancyScore(string, string) (float32, error)

	// CalculateRelevancyScore is the algorithm used to calculate a score relevancy score
	// with considerations made for  latitude and longitude
	calculateRelevancyScoreWithLatLng(string, float32, float32, string) (float32, error)
}

// GetRankForCity is the service the converts the city into a rank between [0, 1]. The algorithm
// to apply in particular is determined at run time by the config
func GetRankForCity(searchTerm, city string, rm rankManager) (float32, error) {
	rank, err := rm.calculateRelevancyScore(searchTerm, city)
	if err != nil {
		l4g.Error("the rank determination algorithm was unable to calculate a rank for the city")
		return 0, err
	}
	return rank, nil
}

// min is a function to calculate min with int types
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
