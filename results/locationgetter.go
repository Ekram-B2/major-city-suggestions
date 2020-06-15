package results

// LatFinder are operations applied to get latitude of a datapoint.
type LatFinder func(dp DataPoint) string

// LongFinder are operations applied to get latitude of a datapoint.
type LongFinder func(dp DataPoint) string

// FindCityLatitude returns lattitude for a data point
func FindCityLatitude(dp DataPoint) string {
	castedDP := dp.(city)
	return castedDP.Lat
}

// FindCityLongitude returns longitude for a data point
func FindCityLongitude(dp DataPoint) string {
	castedDP := dp.(city)
	return castedDP.Long
}

// GetLongitudeForDataPoint is a factory that returns the longitude finder algorithm based on the datapoint
func GetLongitudeForDataPoint(dataPoint string) LongFinder {
	switch dataPoint {
	case "city":
		return FindCityLongitude
	default:
		return FindCityLongitude
	}
}

// GetLatitudeForDataPoint is a factory that returns the latitude finder algorithm based on the datapoint
func GetLatitudeForDataPoint(dataPoint string) LatFinder {
	switch dataPoint {
	case dataPoint:
		return FindCityLatitude
	default:
		return FindCityLatitude
	}
}
