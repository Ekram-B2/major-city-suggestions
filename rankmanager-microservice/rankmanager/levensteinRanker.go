package rankmanager

// LevenSteinRanker is a ranker whose algorithms depend on the levenstein edit formula
type levenSteinRanker struct{}

// CalculateRelevancyScore is the algorithm used to calculate a score
func (lr levenSteinRanker) calculateRelevancyScore(searchTerm string, city string) (float32, error) {
	var score float32
	// 1. Calcuate distance with just the characters
	score += float32(lr.calculateDistanceWithCharacters(searchTerm, city))
	// 2. Calculate score using latitude and longitude measurements that is the need

	// 3. Return score
	return score, nil
}

// calculateDistanceWithCharacters is an unexported function applied to the min distance
// using just the characters
func (lr *levenSteinRanker) calculateDistanceWithCharacters(searchTerm, city string) float32 {
	// apply levenstein distance
	if len(searchTerm) == 0 {
		return float32(len(city))
	}

	if len(city) == 0 {
		return float32(len(searchTerm))
	}

	matrix := make([][]int, len(searchTerm)+1)

	for i := 0; i < len(searchTerm)+1; i++ {
		matrix[i] = make([]int, len(city)+1)
	}

	for i := 1; i < len(searchTerm)+1; i++ {
		matrix[i][0] = matrix[i-1][0] + 1
	}

	for i := 1; i < len(city)+1; i++ {
		matrix[0][i] = matrix[0][i-1] + 1
	}

	for i := 1; i < len(searchTerm)+1; i++ {
		for j := 1; j < len(city)+1; j++ {
			if searchTerm[i-1] == city[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = 1 + min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1]))
			}
		}
	}

	var longerTerm string
	if len(searchTerm) > len(city) {
		longerTerm = searchTerm
	} else {
		longerTerm = city
	}

	return float32(matrix[len(searchTerm)][len(city)]) / float32(len(longerTerm))
}

func (lr levenSteinRanker) calculateRelevancyScoreWithLatLng(string, float32, float32, string) (float32, error) {
	return 0, nil
}
