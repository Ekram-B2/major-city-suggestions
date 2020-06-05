package results

type Cities struct {
	Members         []datapoints.dataPoint
	containsMembers bool
}

// GetSource return non-verbose information about the store
func (c Cities) GetSource() string {
	return "FILE"
}

// GetVerboseSource returns verbose information about the store
func (c Cities) GetVerboseSource() string {
	return "The data source is a file stored on the hosting server"
}

// CombineWith takes two results and performs a join operation
func (c Cities) CombineWith(r Results) {
	c.Members = append(c.Members, r.GetView()...)
}

// AddDataPoint takes two results and performs a join operation
func (c Cities) AddDataPoint(d dataPoint) {
	c.Members = append(c.Members, d)
}

// GetView presents a view of the Results within a linear data structure
func (c Cities) GetView() []dataPoint {
	return c.Members
}
