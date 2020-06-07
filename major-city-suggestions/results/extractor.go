package results

import "errors"

// extractor are operations applied to get the parts of the dataset corresponding to the relevant datapoint
type extractor func(map[string]interface{}) (interface{}, error)

// GetCitySetFromDataset is an implementation to extract cities from the dataset
func GetCitySetFromDataset(dataSet map[string]interface{}) (interface{}, error) {
	var empty interface{}

	if _, ok := dataSet["cities"]; !ok {
		return dataSet["cities"], nil
	}

	return empty, errors.New("unable to located cities")
}

// GetExtractorForDataPoint is a generator applied to acquire the extractor operation for the data point
func GetExtractorForDataPoint(dataPoint string) extractor {
	switch dataPoint {
	case "city":
		return GetCitySetFromDataset
	default:
		return GetCitySetFromDataset
	}
}
