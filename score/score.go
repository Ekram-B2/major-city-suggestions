package score

import (
	"github.com/major-city-suggestions/datastore"
)

// CalculateRelevancyScore is the algorithm used to calculate a score
func CalculateRelevancyScore(searchTerm string, city datastore.LargeCity) (float32, error) {
	var score float32
	// 1. Calcuate distance with just the characters
	score += float32(calculateDistanceWithCharacters(searchTerm, city.City))
	// 2. Calculate score using latitude and longitude measurements that is the need

	// 3. Return score
	return score, nil
}

// calculateDistanceWithCharacters is an unexported function applied to the min distance
// using just the characters
func calculateDistanceWithCharacters(searchTerm, city string) int {
	// apply levenstein distance
	if len(searchTerm) == 0 {
		return len(city)
	}

	if len(city) == 0 {
		return len(searchTerm)
	}

	matrix := make([][]int, len(searchTerm))
	for i := 0; i < len(city); i++ {
		matrix[i] = make([]int, len(city))
	}
	for i := 0; i < len(searchTerm); i++ {
		matrix[0][i] = i
	}

	for i := 0; i < len(city); i++ {
		matrix[i][0] = i
	}

	for i := 1; i < len(searchTerm); i++ {
		for j := 1; j < len(city); j++ {
			if searchTerm[i] == city[j] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = 1 + min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1]))
			}
		}
	}
	return matrix[len(searchTerm)][len(city)]
}

// min is a function to calculate min with int types (the built-in tool)
// is set up for float64
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
