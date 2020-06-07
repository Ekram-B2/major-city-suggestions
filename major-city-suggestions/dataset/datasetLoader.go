package dataset

import (
	"encoding/json"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"
)

// DataLoader is a type referencing operations used to load file paths that make up a dataset based on file type
type DataLoader func() (map[string][]string, error)

// LoadPersistanceFiles gets a map of persistant files based on file type
func LoadPersistanceFiles(manifestPathGetter manifestPathGetter, dataSetBuilder dataSetBuilder) (map[string][]string, error) {

	var manifest manifest

	// 1. Open the file storing the dataset manifest
	datasetBuffer, err := os.Open(manifestPathGetter())

	if err != nil {

		l4g.Error("unable to open the manifest file: %s", err.Error())
		return nil, err
	}

	defer datasetBuffer.Close()

	// 2. Extract a byte stream from the datasetBuffer
	byteStream, err := ioutil.ReadAll(datasetBuffer)
	if err != nil {
		l4g.Error("unable to load byte stream from provided file %s", err.Error())
		return nil, err
	}

	// 3. Unmarshall the byte stream to fit a go structure representation
	err = json.Unmarshal(byteStream, &manifest)
	if err != nil {
		l4g.Error("unable to unmarshall byte stream into data state structure %s", err.Error())
		return nil, err
	}
	// 4. Build dataset from the files stored within the manifest
	dataset := dataSetBuilder(manifest)
	// 5. return dataset and a nil error
	return dataset, nil
}
