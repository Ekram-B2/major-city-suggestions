package config

import (
	"encoding/json"
	"io/ioutil"
	"os"

	l4g "github.com/alecthomas/log4go"
)

// LoadConfiguration used to load the configuration information to a go structure
func LoadConfiguration(getConfigPointer getConfigPointer) (config Config, err error) {
	// 1. Open the file storing the config information that we transform
	configurationBuff, err := os.Open(getConfigPointer())

	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to open the provided file: %s", err.Error())
		return Config{}, err
	}

	defer configurationBuff.Close()

	// 2. Extract a byte stream from the configurationBuff
	byteStream, err := ioutil.ReadAll(configurationBuff)
	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to load byte stream from provided file: %s", err.Error())
		// return an empty implementation paired with an error
		return Config{}, err
	}
	// 3. Unmarshall the byte stream to fit a go structure representation
	err = json.Unmarshal(byteStream, &config)
	if err != nil {
		l4g.Error("OPERATION-ERROR: unable to unmarshall byte stream into data state structure: %s", err.Error())
		return Config{}, err
	}

	// 4. return a loaded configuration
	return config, nil
}
