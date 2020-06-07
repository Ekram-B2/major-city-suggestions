package config

import "os"

// configPathGetter defined operations whose logic is to return a file path to the configuration object
type configPathGetter func() string

// Applied to prevent use of magic numbers
var developmentEnvironmentVar string = "DEV"
var productionEnvironmentVar string = "PROD"
var productionEnvironment string = "0"
var developementEnvironment string = "1"

// getDefaultConfigPath is an implementation applied to return a file path to a configuration object
func getDefaultConfigPath() string {

	// 1. Determine whether the environment is development or production
	env := os.Getenv("DEVELOPMENT")

	// 2. Given the run time environment, determine the path which stores the config information
	if env == productionEnvironment {
		return os.Getenv(productionEnvironmentVar)
	}
	// If it is not the production environment, then it is the developement environment we work within
	return os.Getenv(developmentEnvironmentVar)

}

// GetConfigPathOp is a factory applied to the operation for getting the operation which will be
// applied to determine the path to a configuration object
func GetConfigPathOp(opType string) configPathGetter {
	switch opType {
	case "default":
		return getDefaultConfigPath
	default:
		// the default path is getDefaultConfigPath as no other implemention is presently supported
		return getDefaultConfigPath
	}
}
