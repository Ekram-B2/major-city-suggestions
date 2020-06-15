# dataset
--
    import "github.com/Ekram-B2/suggestionsmanager/dataset"


## Usage

#### func  LoadDataSetFilesDefault

```go
func LoadDataSetFilesDefault(config config.Config, buildDataSetRepresentation DataSetBuilder) ([]string, error)
```
LoadDataSetFilesDefault gets a map of files in local persistance based on file
type

#### type DataSetBuilder

```go
type DataSetBuilder func(manifest Manifest) []string
```

DataSetBuilder defines the operation for converting a manifest into a dataset

#### func  GetDatasetBuilder

```go
func GetDatasetBuilder(opType string) DataSetBuilder
```
GetDatasetBuilder is a factory that returns the dataSetBuilder op to apply which
builds files out to match a format consistent with some designated schema

#### type DataSetLoader

```go
type DataSetLoader func(config.Config, DataSetBuilder) ([]string, error)
```

DataSetLoader is a type defining operations used to load file paths that make up
a dataset

#### func  GetDataSetLoader

```go
func GetDataSetLoader(opType string) DataSetLoader
```
GetDataSetLoader is a factory applied to get the dataLoaderOperation given the
opType

#### type Manifest

```go
type Manifest interface {
	// getView is applied standardize the client representation of the files
	GetView() []string
}
```

Manifest is a type that stores the list of file paths making up a dataset
