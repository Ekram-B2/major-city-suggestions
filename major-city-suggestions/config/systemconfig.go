package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"
)

var localEnvironmentPath string = "LOC"
var productionEnvironmentPath string = "PROD"

// SystemConfig is the config info set within production and development modes
type SystemConfig struct {
	DataPointType string `json:"datapointtype"`
	FileType      string `json:"filetype"`
}

// LoadConfiguration used to load the configuration information to a go structure
func (sc SystemConfig) LoadConfiguration() (Config, error) {

	// 1. Determine if the config file is for the development or production build
	path := getConfigPath()

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
	err = json.Unmarshal(byteStream, sc)
	if err != nil {
		l4g.Error("unable to unmarshall byte stream into data state structure %s", err.Error())
		return SystemConfig{}, err
	}

	// 5. nil for no error
	return sc, nil
}

func getConfigPath() string {
	// 1. Determine whether the environment is local or production
	env := os.Getenv("LOCAL")

	// 2. Given the run time environment, determine the path which stores the config information
	if env == "1" {
		return os.Getenv(localEnvironmentPath)
	}
	return os.Getenv(productionEnvironmentPath)

}
