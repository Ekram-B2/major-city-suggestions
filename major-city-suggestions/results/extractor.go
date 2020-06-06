package results

import "errors"

type extractor func(map[string]interface{}) (interface{}, error)

func GetSampleSetFromDataset(dataSet map[string]interface{}) (interface{}, error) {
	var empty interface{}

	if _, ok := dataSet["cities"]; !ok {
		return dataSet["cities"], nil
	}

	return empty, errors.New("unable to located cities")
}
