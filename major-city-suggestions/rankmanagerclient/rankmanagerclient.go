package rankmanagerclient

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	l4g "github.com/alecthomas/log4go"

	"github.com/major-city-suggestions/major-city-suggestions/config"
)

// RankManagerClient is the client software through which to commit requests to the micro-service
type RankManagerClient struct {
	// Part of the representation of the Rank Manager Client is the configuration
	Config config.SystemConfiguration
}

// Rank is the definition of the score computed by the microservice
type Rank struct {
	CityName string  `json:"Cityname"`
	Rank     float32 `json:"Rank"`
}

// GetRank is the algorithm used to retreive rank for the city name
func (client *RankManagerClient) GetRank(searchTerm, city string) (*Rank, error) {
	// 1. Commit Get request to retreive rank from the remote server
	resp, err := http.Get(getURLToRankMicroService(searchTerm, city))
	if err != nil {
		l4g.Error("Unable to process request for rank from remote server: $s", err.Error())
		return nil, err
	}
	defer resp.Body.Close()
	// 2. From the returned body, read the contents of the body out to a stream
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l4g.Error("Unable to read byte stream from response body")
		return nil, err
	}
	// 3. Unmarshall the content of the body out to a go structure and then return that
	var rank Rank

	err = json.Unmarshal(respBody, &rank)
	if err != nil {
		//l4g.Error("Unable to unmarshall byte stream into rank structure %s", err.Error())
		return nil, err
	}

	// 4. Return the rank
	return &rank, nil
}

func getURLToRankMicroService(searchTerm, city string) string {
	modifiedSearchTerm := strings.ReplaceAll(searchTerm, " ", "%20")
	modifiedCity := strings.ReplaceAll(city, " ", "%20")
	return fmt.Sprintf("http://127.0.0.1:8081/determineRank?searchTerm=%s&city=%s", modifiedSearchTerm, modifiedCity)

}
