package results

// Cities is an implementation of Result
type Cities struct {
	Members         []DataPoint
	containsMembers bool
}

// ContainsMembers retuns the state of whether members are present within the result or not
func (c Cities) ContainsMembers() bool {
	return c.containsMembers == true
}

// CombineWith takes two results and performs a join operation
func (c Cities) CombineWith(r Results) Results {
	c.Members = append(c.Members, r.GetView()...)
	c.containsMembers = true
	return c
}

// AddDataPoint takes two results and performs a join operation
func (c Cities) AddDataPoint(d DataPoint) Results {
	c.Members = append(c.Members, d)
	c.containsMembers = true
	return c
}

// GetView presents a view of the Results within a linear data structure
func (c Cities) GetView() []DataPoint {
	return c.Members
}
