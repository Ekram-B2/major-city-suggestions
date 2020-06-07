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
func NewResultsParser(dataProperties []string, dataPoint string) resultsParser {
	return resultsParser{dataProperties: dataProperties, dataPoint: dataPoint}
}
func (rp resultsParser) verifyIfDataPoint(sample interface{}, dataPoint DataPoint, converter converter) (isDataPoint bool, dp DataPoint) {
	// 1. Get the data point and the subset of properties found to match the minimal set defined within the config

	minimalProperties, dataPoint := converter(sample, dataPoint, rp.dataProperties)

	// 2. See if there are enough properties defined to complete a datapoint sufficient for the remainder of the implementation
	if dataPoint.CanBeCreatedFrom(minimalProperties) != true {
		return false, dataPoint
	}

	// 3. Return datapoint
	return true, dataPoint
}

func (rp resultsParser) ParseUnstructuredResult(dataSet map[string]interface{}, extractor extractor, converter converter, dataPointType string) (res Results, isParsed bool) {

	structuredResultsContainer := GetStructuredResultFormat(dataPointType)

	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("formatting for file does not match the expectations of the parser"))
		}
	}()

	sampleSet, err := extractor(dataSet)

	if err != nil {
		l4g.Error(fmt.Sprintf("unable to extract city related samples from the data store: %s", err.Error()))
		return structuredResultsContainer, false
	}

	for _, sample := range sampleSet.([]interface{}) {
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
