package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"
)

// SystemConfig is the config info set within production and development modes
type SystemConfig struct {
	DataPointType      string   `json:"datapointtype"`
	FileType           string   `json:"filetype"`
	IsRemote           bool     `json:"isremoteclient"`
	MinimalKeySet      []string `json:"minimalkeyset"`
	CharDistCalculator string   `json:"chardistancecalculator"`
}

// GetCharDistCalculator is applied to get the config for the distance calculator between two strings when
// only considering characters
func (sc SystemConfig) GetCharDistCalculator() string {
	return sc.CharDistCalculator
}

// GetMinimalKeySet is applied to get the minimal key set that is representative of what minimal subset of properties
// must be present for the data point to be recognised
func (sc SystemConfig) GetMinimalKeySet() []string {
	return sc.MinimalKeySet
}

// GetDataPointType is applied to get the data point that will be extracted from the data set. Setting the property is
// expecially important if data points are being drawn from files
func (sc SystemConfig) GetDataPointType() string {
	return sc.DataPointType
}

// GetFileType is applied to return the type of file that are being read from. Different file types can have different
// encodings
func (sc SystemConfig) GetFileType() string {
	return sc.FileType
}

// IsRemoteClient is applied to determine if the reader interface is backed by a remote server implementation, or if
// the data set is representative of files stored on the local system
func (sc SystemConfig) IsRemoteClient() bool {
	return sc.IsRemote == true
}

// LoadConfiguration used to load the configuration information to a go structure
func (sc SystemConfig) LoadConfiguration(getConfigPathOp configPathGetter) (Config, error) {

	// 1. Get the path of the config file
	path := getConfigPathOp()

	// 2. Open the file storing the config information that we transform
	configurationBuffer, err := os.Open(path)

	if err != nil {
		l4g.Error("unable to open the provided file: %s", err.Error())
		return SystemConfig{}, err
	}

	defer configurationBuffer.Close()

	// 3. Extract a byte stream from the configurationBuffer
	byteStream, err := ioutil.ReadAll(configurationBuffer)
	if err != nil {

		l4g.Error("unable to load byte stream from provided file %s", err.Error())
		return SystemConfig{}, err
	}

	// 4. Unmarshall the byte stream to fit a go structure representation
	err = json.Unmarshal(byteStream, &sc)
	if err != nil {
		l4g.Error("unable to unmarshall byte stream into data state structure %s", err.Error())
		return SystemConfig{}, err
	}

	// 5. nil for no error
	return sc, nil
}
