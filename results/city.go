package results

// City is a representation of the city stored within persistance
type city struct {
	Name         string `json:"name"`
	Country      string `json:"country"`
	Lat          string `json:"lat"`
	Long         string `json:"long"`
	containsData bool
}

func (c city) GetProperties() []string {
	return []string{"name", "lat", "long", "country"}
}

// GetStateMutators() returns a map that can be applied to mutate the state of an object
func (c city) GetStateMutators() map[string]mutator {

	return map[string]mutator{"name": c.setName,
		"country": c.setCountry,
		"lat":     c.setLat,
		"long":    c.setLong}
}

// CanBeCreatedFrom is applied to determine if a set of properties owned by an object meets the minimum required to
// represent a datapoint
func (c city) CanBeCreatedFrom(foundProperties []string) bool {
	index := 0
	minimalNumOfProperties := len(c.GetProperties())
	for _, prop := range foundProperties {
		switch prop {
		case "name",
			"lat",
			"long",
			"country":
			index++
		default:
			continue
		}
	}

	if index >= minimalNumOfProperties {
		return true
	}
	return false
}

// GetDataPointType is used to return the type of datum the datapoint is representative of
func (c city) GetDataPointType() string {
	return "city"
}

// GetRelevancyKey returns the relevancy key that is applied within a relvancy detection algorithm
func (c city) GetRelevancyKey() string {
	return c.Name
}

// Hash returns a string representation of the data point
func (c city) GetHash() string {
	return c.Name + ", " + c.Country
}

// Equals is applied to compare two data points for equality
func (c city) Equals(d DataPoint) bool {
	defer func() {
		if r := recover(); r != nil {
			// Empty because we will just return false
		}
	}()
	castedDataPoint := d.(city)
	if c.Name == castedDataPoint.Name && c.Country == castedDataPoint.Country &&
		c.Lat == castedDataPoint.Lat && c.Long == castedDataPoint.Long {
		return true
	}

	return false
}
func (c city) setName(val string) DataPoint {
	c.Name = val
	return c
}

func (c city) setCountry(val string) DataPoint {
	c.Country = val
	return c
}

func (c city) setLat(val string) DataPoint {
	c.Lat = val
	return c
}

func (c city) setLong(val string) DataPoint {
	c.Long = val
	return c
}
