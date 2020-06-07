package dataset

// manifestPathGetter defines the operation for returning a manifest path
type manifestPathGetter func() string

// getManifestPath returns the manifest path which points to the file containing meta information
// about the data set
func getManifestPath() string {
	return "major-city-suggestions/dataset/manifest/manifest.json"
}

// getManifestPathOp a factory applied to generate an the operation to retreive the manifest path
func getManifestPathOp(opType string) manifestPathGetter {
	switch opType {
	case "default":
		return getManifestPath
	default:
		return getManifestPath
	}
}
