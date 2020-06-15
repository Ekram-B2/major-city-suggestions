# config
--
    import "github.com/Ekram-B2/suggestionsmanager/config"


## Usage

#### func  GetConfigPath

```go
func GetConfigPath(opType string) getConfigPointer
```
GetConfigPath is a factory for getting the operation which will be applied to
get the path to a configuration object.

#### func  IsDevelopmentEnvironment

```go
func IsDevelopmentEnvironment(deploymentType string) bool
```
IsDevelopmentEnvironment is applied as a way to determine the deployment type

#### type Config

```go
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
```

Config is the config info set within production and development modes

#### func  LoadConfiguration

```go
func LoadConfiguration(getConfigPointer getConfigPointer) (config Config, err error)
```
LoadConfiguration used to load the configuration information to a go structure
