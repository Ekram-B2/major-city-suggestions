package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"
)

var localEnvironmentPath string = "LOC"
var productionEnvironmentPath string = "PROD"

// SystemConfiguration is a structure that loads config info set within
// production and development config files.
type SystemConfiguration struct {
	// Networking configuration information
	Port int `json:"port"`
	// State management information
	Ranker string `json:"ranker"`
}

// LoadConfiguration is a function used to load the configuration information
// to a go structure
func LoadConfiguration() (*SystemConfiguration, error) {

	// 0. Create container to store set config options
	var config SystemConfiguration

	// 1. Determine if the config file is for the development or production build
	path := getConfigPath()

	// 2. Open the file storing the config information that we transform
	configurationBuffer, err := os.Open(path)

	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		l4g.Error("Unable to open the provided file: %s", err.Error())
		return nil, err
	}

	defer configurationBuffer.Close()

	// 2. Extract a byte stream from the configurationBuffer
	byteStream, err := ioutil.ReadAll(configurationBuffer)
	if err != nil {
		// This is a serious problem and the service isn't able to perform what is intended
		l4g.Error("Unable to load byte stream from provided file %s", err.Error())
		return nil, err
	}

	// 3. Unmarshall the byte stream to fit a go structure representation
	err = json.Unmarshal(byteStream, &config)
	if err != nil {
		l4g.Error("Unable to unmarshall byte stream into data state structure %s", err.Error())
		return nil, err
	}

	// 3. Return the cities
	return &config, nil
}

func getConfigPath() string {
	// 1. Set up variable to store path
	var path string

	// 2. Determine whether the environment is local or production
	env := os.Getenv("LOCAL")

	// 3. Given the run time environment, determine the path which stores the config information
	if env == "1" {
		path = os.Getenv(localEnvironmentPath)
	} else {
		path = os.Getenv(productionEnvironmentPath)
	}
	return path
}
