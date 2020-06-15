package dataset

// Manifest is a type that stores the list of file paths making up a dataset
type Manifest interface {
	// getView is applied standardize the client representation of the files
	GetView() []string
}
