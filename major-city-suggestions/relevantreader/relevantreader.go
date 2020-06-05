package relevantreader

// RelevantReader is an interface for reading in only relevant data from a
// persistant store
type RelevantReader interface {

	// ReadRelevant used to access all relevant search terms
	// The return should be a reference type since empty results are valid
	ReadRelevant(string) (*Results, error)
}
