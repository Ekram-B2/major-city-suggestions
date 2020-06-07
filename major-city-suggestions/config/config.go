package config

import "os"

// Config is representative of what the represents a config object
type Config interface {
	LoadConfiguration(configPathGetter) (Config, error)
	GetDataPointType() string
	GetFileType() string
	IsRemoteClient() bool
	GetMinimalKeySet() []string
	GetCharDistCalculator() string
}

// GetConfiguration is a factory applied to load a configuration object based on the breadth of its used
func GetConfiguration(configType string) (Config, error) {
	switch configType {
	case "project":
		config := SystemConfig{}
		loadedConfig, err := config.LoadConfiguration(getConfigPathOp(os.Getenv("CONFIG_PATH_TYPE")))
		if err != nil {
			return config, err
		}
		return loadedConfig, nil
	default:
		config := SystemConfig{}
		loadedConfig, err := config.LoadConfiguration(getConfigPathOp(os.Getenv("CONFIG_PATH_TYPE")))
		if err != nil {
			return config, err
		}
		return loadedConfig, nil
	}
}
