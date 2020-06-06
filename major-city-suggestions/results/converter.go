package results

import (
	"fmt"

	l4g "github.com/alecthomas/log4go"
)

type converter func(interface{}, DataPoint, []string) ([]string, DataPoint)

func ConvertSampleToDataPoint(sample interface{}, dataPoint DataPoint, dataProperties []string) ([]string, DataPoint) {
	// 1. Get the state mutator
	stateMutator := dataPoint.GetStateMutators()

	// 2. Set up deferred function to handle cases of panic
	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("Formatting for file does not match the expectations of the parser"))
		}
	}()

	// 3. Parse through the unstructured sample to see if any properties match with what we expect
	var minimalProperties []string
	for key, value := range sample.(map[string]string) {
		if isAMember(key, dataProperties) {
			minimalProperties = append(minimalProperties, key)
			stateMutator[key](value)
		}
	}
	return minimalProperties, dataPoint
}
