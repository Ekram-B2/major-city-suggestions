package suggestions

import (
	"fmt"
	"net/http"
)

// HandleRequestForSuggestions is the wrapper for all the logic used to build
// the list of suggestions to return
func HandleRequestForSuggestions(rw http.ResponseWriter, req *http.Request) {
	fmt.Printf("Hello world")
}
