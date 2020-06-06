package config

type Config interface {
	LoadConfiguration() (Config, error)
	GetDataPointType() string
	GetFileType() string
	IsRemoteClient() bool
	GetMinimalKeySet() []string
	GetCharDistCalculator() string
}

func GetConfiguration(configType string) (Config, error) {
	switch configType {
	case "system":
		config := SystemConfig{}
		loadedConfig, err := config.LoadConfiguration()
		if err != nil {
			return config, err
		}
		return loadedConfig, nil
	default:
		config := SystemConfig{}
		loadedConfig, err := config.LoadConfiguration()
		if err != nil {
			return config, err
		}
		return loadedConfig, nil
	}
}
