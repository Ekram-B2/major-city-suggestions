package rankmanager

import (
	"net/http"
)

type responseFormat struct {
	CityName string  `json: "cityname"`
	Rank     float32 `json: "rank"`
	Lat      float32 `json: "lat"`
	Lng      float32 `json: "lng"`
}

// HandleRequestToDetermineRank is the wrapper for all the logic used to build
// the list of suggestions to return
func HandleRequestToDetermineRank(rw http.ResponseWriter, req *http.Request) {
	// 0. Create list to store suggestions

}
