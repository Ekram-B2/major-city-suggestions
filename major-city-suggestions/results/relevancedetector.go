package results

import "strings"

// IsRelevantCity is the baseline algorithm used to determine if a city is relevant or not
func IsRelevantCity(searchTerm string, dp DataPoint) bool {
	castedDP := dp.(city)
	return strings.ContainsAny(searchTerm, castedDP.City)
}
