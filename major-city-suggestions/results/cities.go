package results

// Cities is an implementation of Result
type Cities struct {
	Members         []DataPoint
	containsMembers bool
}

// GetSource return non-verbose information about the store
func (c Cities) GetSource() string {
	return "FILE"
}

// ContainsMembers retuns the state of whether members are present within the result or not
func (c Cities) ContainsMembers() bool {
	return c.containsMembers == true
}

// GetVerboseSource returns verbose information about the store
func (c Cities) GetVerboseSource() string {
	return "The data source is a file stored on the hosting server"
}

// CombineWith takes two results and performs a join operation
func (c Cities) CombineWith(r Results) {
	c.Members = append(c.Members, r.GetView()...)
	c.containsMembers = true
}

// AddDataPoint takes two results and performs a join operation
func (c Cities) AddDataPoint(d DataPoint) {
	c.Members = append(c.Members, d)
	c.containsMembers = true
}

// GetView presents a view of the Results within a linear data structure
func (c Cities) GetView() []DataPoint {
	return c.Members
}
