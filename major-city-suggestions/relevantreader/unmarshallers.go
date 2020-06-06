package relevantreader

import (
	"encoding/json"

	l4g "github.com/alecthomas/log4go"
)

// unmarshallIOIntoStruct generates an unstructured representation of the dataset.
func unmarshallJSONIOIntoStruct(byteStream []byte, resultSet map[string]interface{}) (map[string]interface{}, error) {
	// 1. Define container to unmarshall byte stream
	var results map[string]interface{}
	// 2. Apply algorithm for unmarshalling step
	err := json.Unmarshal(byteStream, &results)
	// 3. Log error if necessary
	if err != nil {
		l4g.Error("unable to unmarshall byte stream into a go structure")
		return results, err
	}
	// 4. return results
	return results, nil
}
