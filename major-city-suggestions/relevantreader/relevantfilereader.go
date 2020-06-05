package relevantreader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	l4g "github.com/alecthomas/log4go"
)

//files := map[string][]string{"json": []string{"major-city-suggestions/datamanager/data/ca,json"} }

// Structure applied as an interface to access file reading services
type relevantFileReader struct {
	fileType string
	files    map[string][]string
}

// NewRelevantFileReader is a constructor to return a valid interface
// from which a caller can apply valid read operations. The presently
// supported types for the reader are: `json`
func NewRelevantFileReader(fileType string) *relevantFileReader {
	if fileType != "json" {
		return nil
	}
	return &relevantFileReader{fileType: fileType}
}

// unmarshallIOIntoStruct applies the appropriate encoder based on the specified file type
func (rr relevantFileReader) unmarshallIOIntoStruct(byteStream []byte, results map[string]interface{}) error {
	var err error
	switch rr.fileType {
	case "json":
		err = json.Unmarshal(byteStream, &results)
	default:
		err = json.Unmarshal(byteStream, &results)
	}
	return err
}

// ReadRelevant is used to access all relevant terms
func (rr relevantFileReader) ReadRelevant(searchTerm string) (Results, error) {

	// 1. Retreive unstructured results from datastore
	unstructuredResults := rr.readAll()

	// 2. Parse unstructured results to produce structured results

	resultParser := getParser("city")

	structuredResultsContainer := getStructuredResult("city")

	for _, file := range unstructuredResults {
		resultsForFile := resultParser.parseResult(file)
		structuredResultsContainer.CombineWith(resultsForFile)
	}

	// 3. Filter away irrelevant items from the DataState
	structuredResultsContainer = rr.filterForRelevantDataPoints(structuredResultsContainer, searchTerm)

	// 3. return the new datastate
	return structuredResultsContainer, nil

}

// readAll loads the entire state over all files tracked within local server
func (rr relevantFileReader) readAll() []map[string]interface{} {
	// 1. Define container to store results
	var allResults []map[string]interface{}

	// 2. Perform extraction step over each file
	for _, filePath := range rr.files[rr.fileType] {
		results, err := rr.readAllInFile(filePath)
		if err != nil {
			l4g.Error(fmt.Sprintf("Unable to read file with path: %s", filePath))
			continue
		}
		allResults = append(allResults, results)
	}
	// 3. Return all of the results
	return allResults
}

// readAllInFile loads the entire state from a file
func (rr relevantFileReader) readAllInFile(filePath string) (map[string]interface{}, error) {

	// 1. create a read only file, or log the inability to create a read only file and halt execution
	fileBuff, err := os.Open(filePath)
	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		l4g.Error(err.Error())
		return nil, err
	}

	defer fileBuff.Close()

	// 2. Extract a utf-8 backed bytestream from the file, or log the inability to extract stram and halt execution
	byteStream, err := ioutil.ReadAll(fileBuff)
	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		l4g.Error(err.Error())
		return nil, err
	}

	// 3. Define unstructured form to store results
	results := map[string]interface{}{}

	// 4. Unmarshall the byte stream to the results
	err = rr.unmarshallIOIntoStruct(byteStream, results)

	// 5. Log inability to unmarshall and halt execution if error produced
	if err != nil {
		l4g.Error(err.Error())
		return nil, err
	}

	// 6. Return the cities
	return results, nil

}

// filterForRelevantDataPoints filters entries away from from the structured inputthat are irrelevant to the search term
func (rr relevantFileReader) filterForRelevantDataPoints(results Results, searchTerm string) Results {
	// 1. Create container to store the entries that are determined to be relevant
	structuredResultsContainer := getStructuredResult("city")
	// 2. Apply algorithm on each entry and if deemed relevant, add it to the relevant entry container
	for _, dataPoint := range results.GetView() {
		if rr.isRelevant(searchTerm, dataPoint) {
			structuredResultsContainer.AddDataPoint(dataPoint)
		}
	}
	// 3. Return the modified data state
	return structuredResultsContainer

}

// // isRelevant is the baseline algorithm used to determine if a city is relevant or not
func (rr relevantFileReader) isRelevant(searchTerm string, dp dataPoint) bool {
	return strings.ContainsAny(searchTerm, dp.getRelevancyKey())
}
