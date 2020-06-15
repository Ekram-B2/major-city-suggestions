package results

import (
	"fmt"

	l4g "github.com/alecthomas/log4go"
)

// dataPointConverter is an operation that determines if an extracted sample is a datapoint
type dataPointConverter func(interface{}, string) ([]string, DataPoint)

// GetDataPointConverter is a factory that determines the converter given the datapoint type
func GetDataPointConverter(opType string) dataPointConverter {
	switch opType {
	case "city":
		return ConvertCitySampleToDataPoint
	default:
		return ConvertCitySampleToDataPoint
	}
}

// ConvertCitySampleToDataPoint is an implementation that converts a sanmple to a datapoint if possible
func ConvertCitySampleToDataPoint(sample interface{}, dataPointType string) ([]string, DataPoint) {

	// 1. Set up deferred function to handle cases of panic
	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("OPERATION-ERROR: formatting for the sample does not match the expectations of the parser"))
		}
	}()
	// 2. Parse through the unstructured sample to see if any properties match with what we expect
	minimalProperties := make([]string, 0)
	dataPoint := GetDataPoint(dataPointType)

	for key, value := range sample.(map[string]interface{}) {
		if isAMember(key, dataPoint.GetProperties()) {
			minimalProperties = append(minimalProperties, key)
			value := value.(string)
			stateMutator := dataPoint.GetStateMutators()
			dataPoint = stateMutator[key](value)
		}
	}
	// 3. Return output
	return minimalProperties, dataPoint
}
