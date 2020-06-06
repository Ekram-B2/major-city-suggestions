package rankmanager

type charDistanceCalculator func(string, string) float32

func generateDistanceRanker(rankerType string) charDistanceCalculator {
	switch rankerType {
	case "levenstein":
		return calculateLevensteinDistance
	default:
		return calculateLevensteinDistance
	}
}

// CalculateLevensteinDistance is applied to find a distance using just the characters
func calculateLevensteinDistance(searchTerm, realTerm string) float32 {
	// apply levenstein distance algorithm
	if len(searchTerm) == 0 {
		return float32(len(realTerm))
	}

	if len(realTerm) == 0 {
		return float32(len(searchTerm))
	}

	matrix := make([][]int, len(searchTerm)+1)

	for i := 0; i < len(searchTerm)+1; i++ {
		matrix[i] = make([]int, len(realTerm)+1)
	}

	for i := 1; i < len(searchTerm)+1; i++ {
		matrix[i][0] = matrix[i-1][0] + 1
	}

	for i := 1; i < len(realTerm)+1; i++ {
		matrix[0][i] = matrix[0][i-1] + 1
	}

	for i := 1; i < len(searchTerm)+1; i++ {
		for j := 1; j < len(realTerm)+1; j++ {
			if searchTerm[i-1] == realTerm[j-1] {
				matrix[i][j] = matrix[i-1][j-1]
			} else {
				matrix[i][j] = 1 + min(matrix[i-1][j], min(matrix[i][j-1], matrix[i-1][j-1]))
			}
		}
	}

	var longerTerm string
	if len(searchTerm) > len(realTerm) {
		longerTerm = searchTerm
	} else {
		longerTerm = realTerm
	}

	return float32(matrix[len(searchTerm)][len(realTerm)]) / float32(len(longerTerm))
}

// min is a function to calculate min with int types
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
