package results

// City is a structural representation of the city stored within persistance
type city struct {
	City             string `json:"city"`
	Admin            string `json:"admin"`
	Country          string `json:"country"`
	PopulationProper string `json:"population_proper"`
	ISO2             string `json:"iso2"`
	Capital          string `json:"capital"`
	Lat              string `json:"lat"`
	Lng              string `json:"lng"`
	Population       string `json:"population"`
	containsData     bool
}

// GetStateMutators() ...
func (c city) GetStateMutators() map[string]mutator {

	return map[string]mutator{"city": c.setCity,
		"admin":            c.setAdmin,
		"country":          c.setCountry,
		"population":       c.setPopulation,
		"populationProper": c.setPopulationProper,
		"capital":          c.setCapital,
		"iso2":             c.setISO2,
		"lat":              c.setLat,
		"lng":              c.setLng}
}

// CanBeCreatedFrom is applied to determine if a set of properties owned by an object meets the minimum required to
// represent a datapoint
func (c city) CanBeCreatedFrom(foundProperties []string) bool {
	index := 0
	minimalNumOfProperties := 5
	for _, prop := range foundProperties {
		switch prop {
		case "city",
			"country",
			"iso2",
			"lat",
			"lng":
			index++
		default:
			continue
		}
	}

	if index == minimalNumOfProperties {
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
	return c.City
}

// Hash returns a string representation of the data point
func (c city) GetHash() string {
	return c.City + ", " + c.Admin + ", " + c.ISO2
}

// Equals is applied to compare two data points for equality
func (c city) Equals(d DataPoint) bool {
	defer func() {
		if r := recover(); r != nil {
			// Empty because we will just return false
		}
	}()
	castedDataPoint := d.(city)
	if c.City == castedDataPoint.City && c.Admin == castedDataPoint.Admin && c.Country == castedDataPoint.Country &&
		c.PopulationProper == castedDataPoint.PopulationProper && c.Population == castedDataPoint.Population &&
		c.Lat == castedDataPoint.Lat && c.Lng == castedDataPoint.Lng && c.ISO2 == castedDataPoint.ISO2 &&
		c.Capital == castedDataPoint.Capital {
		return true
	}

	return false
}
func (c city) setCity(val string) DataPoint {
	c.City = val
	return c
}

func (c city) setAdmin(val string) DataPoint {
	c.Admin = val
	return c
}

func (c city) setCountry(val string) DataPoint {
	c.Country = val
	return c
}

func (c city) setPopulationProper(val string) DataPoint {
	c.PopulationProper = val
	return c
}

func (c city) setISO2(val string) DataPoint {
	c.ISO2 = val
	return c
}

func (c city) setCapital(val string) DataPoint {
	c.Capital = val
	return c
}

func (c city) setLat(val string) DataPoint {
	c.Lat = val
	return c
}

func (c city) setLng(val string) DataPoint {
	c.Lng = val
	return c
}

func (c city) setPopulation(val string) DataPoint {
	c.Population = val
	return c
}
