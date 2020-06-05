package results



import (
	"errors"
	"fmt"

	l4g "github.com/alecthomas/log4go"
)

type cityParser struct {
	keys     []string
	filePath string
}

func (cp cityParser) parseSampleInResult(filePath string, dataset interface{}, properties []string) *City {
	var result City

	stateMutators := map[string]func(string){"city": result.setCity,
		"admin":            result.setAdmin,
		"country":          result.setCountry,
		"population":       result.setPopulation,
		"populationProper": result.setPopulationProper,
		"capital":          result.setCapital,
		"ISO2":             result.setISO2,
		"lat":              result.setLat,
		"lng":              result.setLng}

	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("Formatting for file %s does not match the expectations of the parser", filePath))
		}
	}()

	for key, value := range dataset.(map[string]string) {
		if isAMember(key, properties) {
			result.containsData = true
			stateMutators[key](value)
		}
	}

	if result.containsData != true {
		return nil
	}

	return &result
}

func (cp cityParser) parseResult(dataSet map[string]interface{}) Results {
	var cities Cities

	defer func() {
		if r := recover(); r != nil {
			l4g.Error(fmt.Sprintf("formatting for file %s does not match the expectations of the parser", cp.filePath))
		}
	}()

	citySamples, err := getCitiesFromDataset(dataSet)
	if err != nil {
		l4g.Error("unable to extract city related samples from the data store")
		return cities
	}

	for _, sample := range citySamples.([]interface{}) {
		city := cp.parseSampleInResult(cp.filePath, sample, cp.keys)
		if city != nil {
			cities.containsMembers = true
			cities.AddDataPoint(*city)
		}
	}
	return cities
}

func isAMember(key string, properties []string) bool {
	for _, property := range properties {
		if key == property {
			return true
		}
	}
	return false
}

func getCitiesFromDataset(dataSet map[string]interface{}) (interface{}, error) {
	var empty interface{}

	if _, ok := dataSet["cities"]; !ok {
		return dataSet["cities"], nil
	}

	return empty, errors.New("unable to located cities")
}
