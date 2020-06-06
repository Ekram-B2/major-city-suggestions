package rankmanager

import l4g "github.com/alecthomas/log4go"

// rankCalculator is applied to calculate rank of real terms
type rankCalculator interface {
	// calculateRank is the algorithm used to calculate a relevancy score
	calculateRank(string, string) (float32, error)

	// calculateRankWithLatLng is the algorithm used to calculate a score relevancy score with considerations made for latitude and longitude
	calculateRelevancyScoreWithLatLng(string, float32, float32, string) (float32, error)
}

// getRankForRealTerm is the service the converts a real term into a number between [0, 1]. ,s
func getRankForRealTerm(searchTerm, realTerm string, rm rankCalculator) (float32, error) {
	rank, err := rm.calculateRank(searchTerm, realTerm)
	if err != nil {
		l4g.Error("the rank calculator algorithm was unable to get a rank for the city: %s", err.Error())
		return 0, err
	}
	return rank, nil
}

// generateRanker is a factory pattern that gets a concrete rankCalculator
func generateRanker(ranker string) rankCalculator {
	if ranker == "levenstein" {
		return levenSteinRanker{}
	}
	return nil
}

// min is a function to calculate min with int types
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
