package config

// Config is the config info set within production and development modes
type Config struct {
	DataPointType      string `json:"datapointtype"`
	IsRemote           bool   `json:"isremoteclient"`
	DataSetBuildType   string `json:"datasetbuildertype"`
	DataSetLoaderType  string `json:"datasetloadertype"`
	SorterType         string `json:"sortertype"`
	ManifestPath       string `json:"manifestpath"`
	DistanceRankerType string `json:"distancerankertype"`
	LatLngDistCalcType string `json:"latlngdistcalctype"`
	Normalizer         string `json:"normalizertype"`
	CacheKeyType       string `json:"cachekeytype"`
	ByteEncoderType    string `json:"byteencodertype"`
	ByteDecoderType    string `json:"bytedecodertype"`
	CacheType          string `json:"cachetype"`
}

// IsDevelopmentEnvironment is applied as a way to determine the deployment type
func IsDevelopmentEnvironment(deploymentType string) bool {
	if deploymentType == "1" {
		return true
	}
	return false
}
