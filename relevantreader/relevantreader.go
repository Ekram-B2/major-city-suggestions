package relevantreader

import (
	"github.com/Ekram-B2/suggestionsmanager/config"
	"github.com/Ekram-B2/suggestionsmanager/dataset"
	"github.com/Ekram-B2/suggestionsmanager/results"
)

// RelevantReader supports reading relevant data from a persistant store. Relevant data is
// partial segment of the global data set with which a rank can be attributed
type RelevantReader interface {

	// ReadRelevant used to read in relevant data from a persistant store
	ReadRelevant(string) (results.Results, error)
}

// GetReader is a factory applied to get a relevant reader based on the configuraion options
func GetReader(config config.Config) RelevantReader {
	if !config.IsRemote {
		switch config.ReaderType {
		case "json":
			return NewRelevantFileReader(config,
				dataset.GetDatasetBuilder(config.DataSetBuildType),
				dataset.GetDataSetLoader(config.DataSetLoaderType))
		default:
			return NewRelevantFileReader(config,
				dataset.GetDatasetBuilder(config.DataSetBuildType),
				dataset.GetDataSetLoader(config.DataSetLoaderType))
		}

	}
	// This would nominally be the case where a reader would be created to support access to remote clients (e.g. sqldb) but this
	// implementation presently doesn't support that so we return the same result as the local client for now
	switch config.ReaderType {
	case "json":
		return NewRelevantFileReader(config,
			dataset.GetDatasetBuilder(config.DataSetBuildType),
			dataset.GetDataSetLoader(config.DataSetLoaderType))
	default:
		return NewRelevantFileReader(config,
			dataset.GetDatasetBuilder(config.DataSetBuildType),
			dataset.GetDataSetLoader(config.DataSetLoaderType))
	}
}
