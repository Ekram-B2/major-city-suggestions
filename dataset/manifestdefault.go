package dataset

// manifestDefault is an impelmentation of a manifest
type manifestDefault struct {
	Files []string `json:"files"`
}

// GetView is applied standardize the client representation of the files
func (md manifestDefault) GetView() []string {
	return md.Files
}
