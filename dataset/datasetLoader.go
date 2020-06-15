package dataset

import (
	"encoding/json"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"

	"github.com/Ekram-B2/suggestionsmanager/config"
)

// DataSetLoader is a type defining operations used to load file paths that make up a dataset
type DataSetLoader func(config.Config, DataSetBuilder) ([]string, error)

// GetDataSetLoader is a factory applied to get the dataLoaderOperation given the opType
func GetDataSetLoader(opType string) DataSetLoader {
	switch opType {
	case "default":
		return LoadDataSetFilesDefault
	default:
		return LoadDataSetFilesDefault
	}
}

// LoadDataSetFilesDefault gets a map of files in local persistance based on file type
func LoadDataSetFilesDefault(config config.Config, buildDataSetRepresentation DataSetBuilder) ([]string, error) {

	var manifest manifestDefault
	// 1. Open the file storing the dataset manifest
	manifestBuffer, err := os.Open(config.ManifestPath)

	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to open the manifest file: %s", config.ManifestPath)
		return nil, err
	}

	defer manifestBuffer.Close()

	// 2. Extract a byte stream from the datasetBuffer
	byteStream, err := ioutil.ReadAll(manifestBuffer)
	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to load byte stream from provided file: %s", err.Error())
		return nil, err
	}

	// 3. Unmarshall the byte stream to fit a go structure representation
	err = json.Unmarshal(byteStream, &manifest)
	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to unmarshall byte stream into data state structure: %s", err.Error())
		return nil, err
	}

	// 4. Build dataset from the files stored within the manifest
	dataset := buildDataSetRepresentation(manifest)
	// 5. return dataset and a nil error
	return dataset, nil
}
