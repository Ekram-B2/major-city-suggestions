package relevantreader

import (
	"fmt"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"

	"github.com/major-city-suggestions/major-city-suggestions/dataset"
	"github.com/major-city-suggestions/major-city-suggestions/results"
)

// Unexported reader for reading persistant stores encoded within files
type relevantFileReader struct {
	// this is the type of file that the reader can extract from
	fileType string
	// this is a map of all the files that make up the dataset
	dataset map[string][]string
	// this is the datapoint type that is read in
	dataPoint string
}

// NewRelevantFileReader is a constructor used to return a valid reader through which
// valid read operations are applied. The presently supported files types made availible for
// the reader are: `json`
func NewRelevantFileReader(dataPoint, fileType string, dataloader dataset.DataLoader) *relevantFileReader {
	// 1. Resolve case where the file type is not a supported type
	if fileType != "json" {
		return nil
	}
	// 2. Load dataset into project
	dataset, err := dataloader()
	if err != nil {
		l4g.Error("sas unable to read in persistant files from the data set")
		return nil
	}
	// 3. Return a structure provisioned with a specified file type and dataset
	return &relevantFileReader{fileType: fileType, dataset: dataset, dataPoint: dataPoint}
}

// ReadRelevant is applied to return all terms that are deemed relevant to the search term
func (rr relevantFileReader) ReadRelevant(searchTerm string) (results.Results, error) {

	// 1. Retreive unstructured results from datastore
	unstructuredResults := rr.readAll()

	// 2. Get container for structured results
	structuredResultsContainer := results.GetStructuredResultForm(rr.dataPoint)
	// 2. Convert the unstructured results into structured results
	resultsParser := results.NewResultsParser([]string{}, "", "")
	for _, file := range unstructuredResults {
		resultsForFile := resultsParser.ParseUnstructuredResult(file, results.GetSampleSetFromDataset, results.ConvertSampleToDataPoint, rr.dataPoint)
		structuredResultsContainer.CombineWith(resultsForFile)
	}

	// 3. Filter away irrelevant items from the DataState
	structuredResultsContainer = rr.filterForRelevantDataPoints(searchTerm, structuredResultsContainer, results.IsRelevantCity)

	// 4. return the filtered set of results
	return structuredResultsContainer, nil

}

// readAll loads the entire state over all files tracked within local server
func (rr relevantFileReader) readAll() []map[string]interface{} {
	// 1. Define container to store results
	var allResults []map[string]interface{}

	// 2. Perform extraction operation to produce data from an entire dataset
	for _, filePath := range rr.dataset[rr.fileType] {
		results, err := rr.readAllInFile(filePath, unmarshallJSONIOIntoStruct)
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
func (rr relevantFileReader) readAllInFile(filePath string, unmarshallerAlgorithm unmarshaller) (map[string]interface{}, error) {

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
	results, err = unmarshallerAlgorithm(byteStream, results)

	// 5. Log inability to unmarshall and halt execution if error produced
	if err != nil {
		l4g.Error(err.Error())
		return nil, err
	}

	// 6. Return the cities
	return results, nil

}

// filterForRelevantDataPoints filters entries away from from the structured inputthat are irrelevant to the search term
func (rr relevantFileReader) filterForRelevantDataPoints(searchTerm string, resultsSet results.Results, relevanceAlgorithm relevanceDetector) results.Results {
	// 1. Create container to store the entries that are determined to be relevant
	structuredResultsContainer := results.GetStructuredResultForm(rr.dataPoint)
	// 2. Apply algorithm on each entry and if deemed relevant, add it to the relevant entry container
	for _, dataSample := range resultsSet.GetView() {
		if relevanceAlgorithm(searchTerm, dataSample) {
			structuredResultsContainer.AddDataPoint(dataSample)
		}
	}
	// 3. Return the modified data state
	return structuredResultsContainer

}
