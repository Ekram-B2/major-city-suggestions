package relevantreader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	l4g "github.com/alecthomas/log4go"
)

// Unexported reader for reading persistant stores encoded within files
type relevantFileReader struct {
	// this is the type of file that the reader can extract from
	fileType string
	// this is a map of all the files that make up the dataset
	dataset map[string][]string
}

// NewRelevantFileReader is a constructor used to return a valid reader through which
// valid read operations are applied. The presently supported files types made availible for
// the reader are: `json`
func NewRelevantFileReader(fileType, dataset string) *relevantFileReader {
	// 1. Resolve case where the file type is not a supported type
	if fileType != "json" {
		return nil
	}
	// 2. Return a structure provisioned with a specified file type and dataset
	return &relevantFileReader{fileType: fileType, dataset: dataset}
}


// ReadRelevant is applied to return all terms that are deemed relevant to the search term
func (rr relevantFileReader) ReadRelevant(searchTerm string) (Results, error) {

	// 1. Retreive unstructured results from datastore
	unstructuredResults := rr.readAll()

	// 2. Parse unstructured results per file to produce structured results
	resultParser := getParser("city")

	structuredResultsContainer := getStructuredResult("city")

	for _, file := range unstructuredResults {
		resultsForFile := resultParser.parseResult(file)
		structuredResultsContainer.CombineWith(resultsForFile)
	}

	// 3. Filter away irrelevant items from the DataState
	structuredResultsContainer = rr.filterForRelevantDataPoints(structuredResultsContainer, searchTerm)

	// 4. return the filtered set of results
	return structuredResultsContainer, nil

}

// readAll loads the entire state over all files tracked within local server
func (rr relevantFileReader) readAll() []map[string]interface{} {
	// 1. Define container to store results
	var allResults []map[string]interface{}

	// 2. Perform extraction operation to produce data from an entire dataset
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

// readAllInFile loads the entire state from a file. The unmarshaller is the algorithm used to convert 
// a byte stream to fit a go struct
func (rr relevantFileReader) readAllInFile(filePath string, func unmarshaller) (map[string]interface{}, error) {

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
	err = rr.unmarshaller(byteStream, results)

	// 5. Log inability to unmarshall and halt execution if error produced
	if err != nil {
		l4g.Error(err.Error())
		return nil, err
	}

	// 6. Return the cities
	return results, nil

}

// filterForRelevantDataPoints filters entries away from from the structured inputthat are irrelevant to the search term
func (rr relevantFileReader) filterForRelevantDataPoints(results Results, searchTerm string, func relevanceAlgorithm) Results {
	// 1. Create container to store the entries that are determined to be relevant
	structuredResultsContainer := getStructuredResult("city")
	// 2. Apply algorithm on each entry and if deemed relevant, add it to the relevant entry container
	for _, dataPoint := range results.GetView() {
		if relevanceAlgorithm(searchTerm, dataPoint) {
			structuredResultsContainer.AddDataPoint(dataPoint)
		}
	}
	// 3. Return the modified data state
	return structuredResultsContainer

}

// // isRelevant is the baseline algorithm used to determine if a city is relevant or not
func isRelevant(searchTerm string, dp dataPoint) bool {
	return strings.ContainsAny(searchTerm, dp.getRelevancyKey())
}


// unmarshallIOIntoStruct generates an unstructured representation of the dataset.
func unmarshallJSONIOIntoStruct(byteStream []byte, results map[string]interface{}) (Results, error) {
	// 1. Define error type in block scope of the entire function
	var err error
	// 2. Apply encoder based on specified file type
	err = json.Unmarshal(byteStream, &results)
	// 3. return results to caller
	return results, err
}