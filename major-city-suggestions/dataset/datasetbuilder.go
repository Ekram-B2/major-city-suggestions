package dataset

import "strings"

// dataSetBuilder defines the operation for converting a manifest into a dataset
type dataSetBuilder func(manifest manifest) map[string][]string

// manifest is a type that stores the list of file paths making up a dataset based
type manifest struct {
	Files []string `json:"files"`
}

// getDatasetBuilderOp is a factory that returns the dataSetBuilder op to apply
func getDatasetBuilderOp(opType string) dataSetBuilder {
	switch opType {
	case "default":
		return defaultbuildDataSetFrom
	default:
		return defaultbuildDataSetFrom
	}
}

// getExtension the extension of a file from its file path
func getExtension(file string) string {
	sepDelimSlice := strings.Split(file, ".")
	return sepDelimSlice[len(sepDelimSlice)-1]
}

// defaultbuildDataSetFrom is an implementation that builds a dataset from a manifest
func defaultbuildDataSetFrom(manifest manifest) map[string][]string {
	var dataset map[string][]string
	// 1. From the manifest, populate the dataset
	for _, file := range manifest.Files {
		extension := getExtension(file)
		if files, ok := dataset[extension]; ok {
			files = append(files, file)
		} else {
			dataset[extension] = []string{file}
		}
	}
	// 2. Return populated dataset
	return dataset
}
