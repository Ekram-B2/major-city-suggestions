package dataset

// ManifestPathGetter defines the operation for returning a manifest path
type ManifestPathGetter func() string

// getManifestPath returns the manifest path which points to the file containing meta information
// about the data set
func getManifestPath() string {
	return "major-city-suggestions/dataset/manifest/manifest.json"
}

// GetManifestPathOp a factory applied to generate an the operation to retreive the manifest path
func GetManifestPathOp(opType string) ManifestPathGetter {
	switch opType {
	case "default":
		return getManifestPath
	default:
		return getManifestPath
	}
}
