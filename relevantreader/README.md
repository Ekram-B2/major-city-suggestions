# relevantreader
--
    import "github.com/Ekram-B2/suggestionsmanager/relevantreader"


## Usage

#### func  NewRelevantFileReader

```go
func NewRelevantFileReader(config config.Config, dataSetBuilder dataset.DataSetBuilder, dataloader dataset.DataSetLoader) *relevantFileReader
```
NewRelevantFileReader is a constructor used to return a valid reader through
which valid read operations are applied. The presently supported files types
made availible for the reader are: `json`

#### type RelevantReader

```go
type RelevantReader interface {

	// ReadRelevant used to read in relevant data from a persistant store
	ReadRelevant(string) (results.Results, error)
}
```

RelevantReader supports reading relevant data from a persistant store. Relevant
data is partial segment of the global data set with which a rank can be
attributed
