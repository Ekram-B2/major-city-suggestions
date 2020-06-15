package results

import (
	"strings"
)

// RelevanceDetector is applied to define functions which detect whether a term is relevant
type RelevanceDetector func(string, DataPoint) bool

// IsRelevantCity is the baseline algorithm used to determine if a datapoint is relevant or not
func IsRelevantCity(searchTerm string, dp DataPoint) bool {
	return strings.Contains(strings.ToLower(dp.GetHash()), strings.ToLower(searchTerm))
}

// GetRelevanceDetector is a factory applied at run time to get the implementation ofthe relevancy algorithm
func GetRelevanceDetector(dataPoint string) RelevanceDetector {
	switch dataPoint {
	case "city":
		return IsRelevantCity
	default:
		return IsRelevantCity
	}
}
