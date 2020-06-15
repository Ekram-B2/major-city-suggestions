package config

import "os"

// getConfigPointer defined operations whose logic is to return a file path to the configuration object
type getConfigPointer func() string

// getConfigPathDefault is an implementation applied to return a file path pointing to a configuration object
func getConfigPathDefault() string {
	// 1. Given the run time environment, determine the path points to configuration information
	if IsDevelopmentEnvironment(os.Getenv("DEPLOYMENT_TYPE")) {
		return os.Getenv("DEVELOPMENT_CONFIG_PATH")
	}

	// 2. If it is not the production environment, then it is the developement environment we work within
	return os.Getenv("PRODUCTION_CONFIG_PATH")
}

// GetConfigPath is a factory for getting the operation which will be applied to get the path to a configuration object.
func GetConfigPath(opType string) getConfigPointer {
	switch opType {
	case "default":
		return getConfigPathDefault
	default:
		// return the default implementation as no other version is presently supported
		return getConfigPathDefault
	}
}
