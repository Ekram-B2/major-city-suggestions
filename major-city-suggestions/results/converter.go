package results

import (
	"fmt"

	l4g "github.com/alecthomas/log4go"
)

// converter is an operation that determines if an extracted sample is a datapoint
type converter func(interface{}, DataPoint, []string) ([]string, DataPoint)

// ConvertSampleToDataPoint is an implementation that converts a sanmple to a datapoint if possible
func ConvertSampleToDataPoint(sample interface{}, dataPoint DataPoint, dataProperties []string) ([]string, DataPoint) {
	// 1. Get the state mutator

	// 2. Set up deferred function to handle cases of panic
	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("formatting for the sample does not match the expectations of the parser"))
		}
	}()
	// 3. Parse through the unstructured sample to see if any properties match with what we expect
	minimalProperties := make([]string, 0)

	for key, value := range sample.(map[string]interface{}) {
		if isAMember(key, dataProperties) {
			minimalProperties = append(minimalProperties, key)
			value := value.(string)
			stateMutator := dataPoint.GetStateMutators()
			dataPoint = stateMutator[key](value)
		}
	}
	return minimalProperties, dataPoint
}
