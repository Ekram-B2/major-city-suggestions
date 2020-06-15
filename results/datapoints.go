package results

// DataPoint is a representation of a single data unit stored within results
type DataPoint interface {
	// GetDataPointType is applied to determine what type of datapoint is the struct representing
	GetDataPointType() string
	// CanBeCreatedFro( is used to compare a property set against what is minimally required to represent the datapoint
	CanBeCreatedFrom([]string) bool
	// GetStateMutators is used to returned an object that can be applied to mutate the information within the datapoint
	GetStateMutators() map[string]mutator
	// Equals is used to detemine if datapoints are equals
	Equals(DataPoint) bool
	// GetHash is used to return the name of the data point
	GetHash() string
	// GetProperties is used get a keyset of the required properties for the data point
	GetProperties() []string
}

// mutator is a operation that changes the state of a datapoint
type mutator func(string) DataPoint

// GetDataPoint is a factory applied to initialize datapoint types
func GetDataPoint(datapoint string) DataPoint {
	switch datapoint {
	case "city":
		return city{}
	default:
		return city{}
	}
}
