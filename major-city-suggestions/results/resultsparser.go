package results

import (
	"fmt"

	l4g "github.com/alecthomas/log4go"
)

type resultsParser struct {
	dataProperties []string
	filePath       string
	dataPoint      string
}

// NewResultsParser is a constructor to create a resultParser with a valid state
func NewResultsParser(dataProperties []string, filePath, dataPoint string) resultsParser {
	return resultsParser{dataProperties: dataProperties, filePath: filePath, dataPoint: dataPoint}
}
func (rp resultsParser) verifyIfDataPointExists(sample interface{}, converter converter, dataPoint DataPoint) (bool, DataPoint) {

	minimalProperties, dataPoint := converter(sample, dataPoint, rp.dataProperties)

	// 2. See if there are enough properties defined to complete a datapoint sufficient for the remainder of the implementation
	if dataPoint.CanBeCreatedFrom(minimalProperties) != true {
		return false, dataPoint
	}

	// 3. Return datapoint
	return true, dataPoint
}

func (rp resultsParser) ParseUnstructuredResult(dataSet map[string]interface{}, extractor extractor, converter converter, dataPointType string) Results {

	structuredResultsContainer := GetStructuredResultForm(dataPointType)

	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("formatting for file %s does not match the expectations of the parser", rp.filePath))
		}
	}()

	sampleSet, err := extractor(dataSet)

	if err != nil {
		l4g.Error("unable to extract city related samples from the data store")
		return structuredResultsContainer
	}

	for _, sample := range sampleSet.([]interface{}) {
		var datum DataPoint

		isDataPoint, datum := rp.verifyIfDataPointExists(sample, converter, datum)
		if isDataPoint == true {
			structuredResultsContainer.AddDataPoint(datum)
		}
	}

	return structuredResultsContainer
}

func isAMember(key string, properties []string) bool {
	for _, property := range properties {
		if key == property {
			return true
		}
	}
	return false
}
