package results

// Results is an generic interface for the different types of data
type Results interface {
	// GetView returns a generic view of the results data
	GetView() []DataPoint
	// Combine two sets of results
	CombineWith(Results)
	// AddDataPoint adds a new datapoint to the results
	AddDataPoint(DataPoint)
	// ContainsMembers checks to see if there are any members in the Result set
	ContainsMembers() bool
}

// GetStructuredResultFormat is a factory applies to initialize result types
func GetStructuredResultFormat(dataPoint string) Results {
	switch dataPoint {
	case "cities":
		return Cities{}
	default:
		return Cities{}
	}
}
