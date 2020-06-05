package relevantreader

type dataPoint interface {
	// getDataPointType is applied to determine what type of datapoint is the struct representing
	getDataPointType() string
	// getRelevancyKey is applied to determine what the relevance key is of the datapoint
	getRelevancyKey() string
}
