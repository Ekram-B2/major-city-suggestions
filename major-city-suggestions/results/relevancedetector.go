package results

import "strings"

type RelevanceDetector func(string, DataPoint) bool

// IsRelevantCity is the baseline algorithm used to determine if a city is relevant or not
func IsRelevantCity(searchTerm string, dp DataPoint) bool {
	castedDP := dp.(city)
	return strings.ContainsAny(searchTerm, castedDP.City)
}

func GetRelevanceDetector(dataPoint string) RelevanceDetector {
	switch dataPoint {
	case "city":
		return IsRelevantCity
	default:
		return IsRelevantCity
	}
}
