# suggestions
--
    import "github.com/Ekram-B2/suggestionsmanager/suggestions"


## Usage

#### func  HandleRequestForSuggestions

```go
func HandleRequestForSuggestions(rw http.ResponseWriter, req *http.Request)
```
HandleRequestForSuggestions handles the logic used to build the list of
suggestions to return back the the caller

#### type Suggestion

```go
type Suggestion struct {
	Name      string  `json:"name"`
	Latitude  string  `json:"latitude"`
	Longitude string  `json:"longitude"`
	Score     float32 `json:"score"`
}
```

Suggestion is the transformed output presented bacl to the client
