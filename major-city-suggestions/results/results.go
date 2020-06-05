package results

// Results is an generic interface for the different types of data
type Results interface {
	// GetSource returns non-verbose information about the store
	GetSource() string
	// GetVerboseSource returns verbose information about the store
	GetVerboseSource() string
	// GetView returns a generic view of the results data
	GetView() []dataPoint
	// Combine two sets of results
	CombineWith(Results)
	// AddDataPoint adds a new datapoint to the results
	AddDataPoint(dataPoint)
}

func getStructuredResult(dataPoint string) Results {
	switch dataPoint {
	case "city":
		return Cities{}
	default:
		return Cities{}
	}
}
