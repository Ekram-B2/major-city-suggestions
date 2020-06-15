package relevantreader

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"

	"github.com/Ekram-B2/suggestionsmanager/config"
	"github.com/Ekram-B2/suggestionsmanager/dataset"
	"github.com/Ekram-B2/suggestionsmanager/results"
)

// Unexported reader for reading persistant stores encoded within files
type relevantFileReader struct {
	// this is a map of all the files that make up the dataset
	dataset []string
	// this is the datapoint type that is read in
	config config.Config
}

// NewRelevantFileReader is a constructor used to return a valid reader through which
// valid read operations are applied. The presently supported files types made availible for
// the reader are: `json`
func NewRelevantFileReader(config config.Config, dataSetBuilder dataset.DataSetBuilder, dataloader dataset.DataSetLoader) *relevantFileReader {

	// 1. Load dataset into project
	dataset, err := dataloader(config, dataSetBuilder)

	if err != nil {
		l4g.Error("OPERATION-ERROR: was unable to read in persistant files")
		return nil
	}
	// 3. Return a structure provisioned with a specified file type and dataset
	return &relevantFileReader{dataset: dataset, config: config}
}

// ReadRelevant is applied to return all terms that are deemed relevant to the search term
func (rr relevantFileReader) ReadRelevant(searchTerm string) (results.Results, error) {

	// 1. Retreive unstructured results from datastore
	unstructuredResults := rr.readAll()
	// 2. Get container for structured results
	structuredResults := results.GetStructuredResult(rr.config.DataPointType)

	// 3. Convert the unstructured results into structured results
	resultsParser := results.NewResultsParser(rr.config.DataPointType)
	for _, samplesInFile := range unstructuredResults {
		resultsForFile, wasFileParsed := resultsParser.ParseUnstructuredResult(samplesInFile,
			results.GetDataPointConverter(rr.config.DataPointType),
			rr.config.DataPointType)

		if wasFileParsed == true {
			structuredResults = structuredResults.CombineWith(resultsForFile)
		}
	}

	// 4. Filter away irrelevant items from the DataState
	structuredResults = rr.filterForRelevantDataPoints(searchTerm, structuredResults, results.GetRelevanceDetector(rr.config.DataPointType))
	// 5. return the filtered set of results
	return structuredResults, nil

}

// readAll loads the entire state over all files tracked within local server
func (rr relevantFileReader) readAll() []interface{} {
	// 1. Define container to store results
	var allResults []interface{}

	// 2. Perform extraction operation to produce data from an entire dataset
	for _, filePath := range rr.dataset {
		results, err := rr.readAllInFile(filePath)
		if err != nil {
			l4g.Error(fmt.Sprintf("OPERATION-ERROR: unable to read file with path: %s", filePath))
			continue
		}
		allResults = append(allResults, results)
	}
	// 3. Return all of the results
	return allResults
}

// readAllInFile loads the entire state from a file. The unmarshaller is the algorithm used to convert
// a byte stream to fit a go struct
func (rr relevantFileReader) readAllInFile(filePath string) (results interface{}, err error) {

	// 1. create a read only file, or log the inability to create a read only file and halt execution
	fileBuff, err := os.Open(filePath)
	if err != nil {
		l4g.Error(fmt.Sprintf("OPERATION-ERROR: unable to open file: %s", err.Error()))
		return nil, err
	}

	defer fileBuff.Close()

	// 2. Extract a utf-8 backed bytestream from the file, or log the inability to extract stream and halt execution
	byteStream, err := ioutil.ReadAll(fileBuff)

	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to create byte stream from file: %s", err.Error())
		return nil, err
	}
	// 3. Unmarshall the byte stream to the results

	err = json.Unmarshal(byteStream, &results)

	if err != nil {
		l4g.Error(fmt.Sprintf("OPERATION-ERROR: unable to unmarshall bytestream into go object: %s", err.Error()))
		return nil, err
	}

	// 4. Return the cities
	return results, nil

}

// filterForRelevantDataPoints filters entries away from from the structured inputthat are irrelevant to the search term
func (rr relevantFileReader) filterForRelevantDataPoints(searchTerm string, resultsSet results.Results, relevanceAlgorithm results.RelevanceDetector) results.Results {
	// 1. Create container to store the entries that are determined to be relevant
	structuredResultsContainer := results.GetStructuredResult(rr.config.DataPointType)
	// 2. Apply algorithm on each entry and if deemed relevant, add it to the relevant entry container
	for _, dataSample := range resultsSet.GetView() {
		if relevanceAlgorithm(searchTerm, dataSample) {
			structuredResultsContainer = structuredResultsContainer.AddDataPoint(dataSample)
		}
	}

	// 3. Return the modified data state
	return structuredResultsContainer

}
