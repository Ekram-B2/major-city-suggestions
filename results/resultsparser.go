package results

import (
	"fmt"

	l4g "github.com/alecthomas/log4go"
)

type resultsParser struct {
	dataProperties []string
	dataPoint      string
}

// NewResultsParser is a constructor to create a resultParser with a valid state
func NewResultsParser(dataPoint string) resultsParser {
	return resultsParser{dataPoint: dataPoint}
}
func (rp resultsParser) verifyIfDataPoint(sample interface{}, dataPoint DataPoint, converter dataPointConverter) (isDataPoint bool, dp DataPoint) {
	// 1. Get the data point and the subset of properties found to match the minimal set defined within the config

	minimalProperties, dataPoint := converter(sample, rp.dataPoint)

	// 2. See if there are enough properties defined to complete a datapoint sufficient for the remainder of the implementation
	if dataPoint.CanBeCreatedFrom(minimalProperties) != true {
		return false, dataPoint
	}

	// 3. Return datapoint
	return true, dataPoint
}

func (rp resultsParser) ParseUnstructuredResult(samplesInFile interface{}, converter dataPointConverter, dataPointType string) (res Results, isParsed bool) {

	structuredResultsContainer := GetStructuredResult(dataPointType)

	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("OPERATION-ERROR: formatting for file does not match the expectations of the parser"))
		}
	}()
	for _, sample := range samplesInFile.([]interface{}) {
		dataPoint := GetDataPoint(rp.dataPoint)
		isDataPoint, dp := rp.verifyIfDataPoint(sample, dataPoint, converter)
		if isDataPoint == true {
			structuredResultsContainer = structuredResultsContainer.AddDataPoint(dp)
		}
	}

	return structuredResultsContainer, true
}

func isAMember(key string, properties []string) bool {
	// 1. Iterate over property list to determine if key belongs to set
	for _, property := range properties {
		if key == property {
			// return true if key is found in set
			return true
		}
	}
	// return false if key is not found in set
	return false
}
