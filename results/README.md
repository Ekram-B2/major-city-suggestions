# results
--
    import "github.com/Ekram-B2/suggestionsmanager/results"


## Usage

#### func  FindCityLatitude

```go
func FindCityLatitude(dp DataPoint) string
```
FindCityLatitude returns lattitude for a data point

#### func  FindCityLongitude

```go
func FindCityLongitude(dp DataPoint) string
```
FindCityLongitude returns longitude for a data point

#### func  GetDataPointConverter

```go
func GetDataPointConverter(opType string) dataPointConverter
```
GetDataPointConverter is a factory that determines the converter given the
datapoint type

#### func  IsRelevantCity

```go
func IsRelevantCity(searchTerm string, dp DataPoint) bool
```
IsRelevantCity is the baseline algorithm used to determine if a datapoint is
relevant or not

#### func  NewResultsParser

```go
func NewResultsParser(dataPoint string) resultsParser
```
NewResultsParser is a constructor to create a resultParser with a valid state

#### type Cities

```go
type Cities struct {
	Members []DataPoint
}
```

Cities is an implementation of Result

#### func (Cities) AddDataPoint

```go
func (c Cities) AddDataPoint(d DataPoint) Results
```
AddDataPoint takes two results and performs a join operation

#### func (Cities) CombineWith

```go
func (c Cities) CombineWith(r Results) Results
```
CombineWith takes two results and performs a join operation

#### func (Cities) ContainsMembers

```go
func (c Cities) ContainsMembers() bool
```
ContainsMembers retuns the state of whether members are present within the
result or not

#### func (Cities) GetView

```go
func (c Cities) GetView() []DataPoint
```
GetView presents a view of the Results within a linear data structure

#### type DataPoint

```go
type DataPoint interface {
	// GetDataPointType is applied to determine what type of datapoint is the struct representing
	GetDataPointType() string
	// CanBeCreatedFro( is used to compare a property set against what is minimally required to represent the datapoint
	CanBeCreatedFrom([]string) bool
	// GetStateMutators is used to returned an object that can be applied to mutate the information within the datapoint
	GetStateMutators() map[string]mutator
	// Equals is used to detemine if datapoints are equals
	Equals(DataPoint) bool
	// GetHash is used to return the name of the data point
	GetHash() string
	// GetProperties is used get a keyset of the required properties for the data point
	GetProperties() []string
}
```

DataPoint is a representation of a single data unit stored within results

#### func  ConvertCitySampleToDataPoint

```go
func ConvertCitySampleToDataPoint(sample interface{}, dataPointType string) ([]string, DataPoint)
```
ConvertCitySampleToDataPoint is an implementation that converts a sanmple to a
datapoint if possible

#### func  GetDataPoint

```go
func GetDataPoint(datapoint string) DataPoint
```
GetDataPoint is a factory applied to initialize datapoint types

#### type LatFinder

```go
type LatFinder func(dp DataPoint) string
```

LatFinder are operations applied to get latitude of a datapoint.

#### func  GetLatitudeForDataPoint

```go
func GetLatitudeForDataPoint(dataPoint string) LatFinder
```
GetLatitudeForDataPoint is a factory that returns the latitude finder algorithm
based on the datapoint

#### type LongFinder

```go
type LongFinder func(dp DataPoint) string
```

LongFinder are operations applied to get latitude of a datapoint.

#### func  GetLongitudeForDataPoint

```go
func GetLongitudeForDataPoint(dataPoint string) LongFinder
```
GetLongitudeForDataPoint is a factory that returns the longitude finder
algorithm based on the datapoint

#### type RelevanceDetector

```go
type RelevanceDetector func(string, DataPoint) bool
```

RelevanceDetector is applied to define functions which detect whether a term is
relevant

#### func  GetRelevanceDetector

```go
func GetRelevanceDetector(dataPoint string) RelevanceDetector
```
GetRelevanceDetector is a factory applied at run time to get the implementation
ofthe relevancy algorithm

#### type Results

```go
type Results interface {
	// GetView returns a generic view of the results data
	GetView() []DataPoint
	// Combine two sets of results
	CombineWith(Results) Results
	// AddDataPoint adds a new datapoint to the results
	AddDataPoint(DataPoint) Results
	// ContainsMembers checks to see if there are any members in the Result set
	ContainsMembers() bool
}
```

Results is an generic interface for the different types of data

#### func  GetStructuredResult

```go
func GetStructuredResult(dataPoint string) Results
```
GetStructuredResult is a factory applies to initialize result types
