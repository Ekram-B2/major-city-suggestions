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

// RankManagerClient is the client software that can retreive from the micro service managing the ranks
type RankManagerClient struct {
	Config config.SystemConfig
}

// rank is the definition for what is retreived from the microservice
type rank struct {
	Name string  `json:"name"`
	Rank float32 `json:"rank"`
}

// GetRank is the algorithm used to retreive rank for the real term retreived from persistance against the research term
func (client *RankManagerClient) GetRank(searchTerm, realTerm string) (rank, error) {
	// 1. Commit GET request to retreive the rank for this datapoint from the remote server
	resp, err := http.Get(createURL(searchTerm, realTerm))
	if err != nil {
		l4g.Error("unable to process request to retreive the rank: %s", err.Error())
		return rank{}, err
	}
	defer resp.Body.Close()
	// 2. Read the contents of the return out to a stram
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		l4g.Error("Unable to read byte stream from response body: %s", err.Error())
		return rank{}, err
	}
	// 3. Unmarshall the contents out to a rank object
	var returnedRank rank

	err = json.Unmarshal(contents, &returnedRank)
	if err != nil {
		l4g.Error("unable to unmarshall byte stream into rank structure %s", err.Error())
		return rank{}, err
	}

	// 4. Return the rank
	return returnedRank, nil
}

func createURL(searchTerm, realTerm string) string {
	modifiedSearchTerm := strings.ReplaceAll(searchTerm, " ", "%20")
	modifiedRealTerm := strings.ReplaceAll(realTerm, " ", "%20")
	return fmt.Sprintf("http://127.0.0.1:8081/determineRank?searchTerm=%s&realTerm=%s", modifiedSearchTerm, modifiedRealTerm)

}
