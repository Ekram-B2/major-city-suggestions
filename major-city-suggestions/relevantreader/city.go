package relevantreader

// City is a structural representation of the city datums stored within persistance
type City struct {
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

func (c City) getDataPointType() string {
	return "city"
}

func (c City) getRelevancyKey() string {
	return c.City
}

func (c City) setCity(val string) {
	c.City = val
}

func (c City) setAdmin(val string) {
	c.Admin = val
}

func (c City) setCountry(val string) {
	c.Country = val
}

func (c City) setPopulationProper(val string) {
	c.PopulationProper = val
}

func (c City) setISO2(val string) {
	c.ISO2 = val
}

func (c City) setCapital(val string) {
	c.Capital = val
}

func (c City) setLat(val string) {
	c.Lat = val
}

func (c City) setLng(val string) {
	c.Lng = val
}

func (c City) setPopulation(val string) {
	c.Population = val
}
