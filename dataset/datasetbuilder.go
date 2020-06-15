package dataset

import "strings"

// DataSetBuilder defines the operation for converting a manifest into a dataset
type DataSetBuilder func(manifest Manifest) []string

// GetDatasetBuilder is a factory that returns the dataSetBuilder op to apply which
// builds files out to match a format consistent with some designated schema
func GetDatasetBuilder(opType string) DataSetBuilder {
	switch opType {
	case "default":
		return buildDataSetFilesDefault
	default:
		return buildDataSetFilesDefault
	}
}

// createDataSetDefault is an implementation that builds a representation of a dataset
// based on the schema
//		{
//			fileType1: [filesOfKeyType...]
//			fileType2: [filesOfKeyType...]
//			...
//		}
//
// from a manifest with the schema
// 		{
//			files : [files...]
//		}
//
func buildDataSetFilesDefault(manifest Manifest) []string {
	dataset := make([]string, 0)
	// 1. From the manifest, populate the dataset
	for _, file := range manifest.GetView() {
		extension := getExtension(file)
		if extension == "json" {
			dataset = append(dataset, file)
		}
	}
	// 2. Return populated dataset
	return dataset
}

// getExtension returns the extension of a file given its file path
func getExtension(file string) string {
	sepSlice := strings.Split(file, ".")
	return sepSlice[len(sepSlice)-1]
}
