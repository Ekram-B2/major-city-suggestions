package dataset

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"

	l4g "github.com/alecthomas/log4go"
)

type DataLoader func() (map[string][]string, error)

type manifest struct {
	files []string `json:"files"`
}

// LoadPersistanceFiles gets a map of persistant files based ons file type
func LoadPersistanceFiles() (map[string][]string, error) {

	var dataset map[string][]string
	var manifest manifest
	// 1. Open the file storing the dataset manifest
	datasetBuffer, err := os.Open(getManifestPath())

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
	for _, file := range manifest.files {
		extension := getExtension(file)
		if files, ok := dataset[extension]; ok {
			files = append(files, file)
		} else {
			dataset[extension] = []string{file}
		}
	}
	return dataset, nil
}

func getManifestPath() string {
	return "major-city-suggestions/dataset/manifest.json"
}

func getExtension(file string) string {
	sepDelimSlice := strings.Split(file, ".")
	return sepDelimSlice[len(sepDelimSlice)-1]
}
