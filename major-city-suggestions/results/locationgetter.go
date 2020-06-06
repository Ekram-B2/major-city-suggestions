package results

type LatFinder func(dp DataPoint) string

type LngFinder func(dp DataPoint) string

func FindCityLatitude(dp DataPoint) string {
	castedDP := dp.(city)
	return castedDP.Lat
}

func FindCityLongitude(dp DataPoint) string {
	castedDP := dp.(city)
	return castedDP.Lng
}

func GetLongitudeForDataPoint(dataPoint string) LngFinder {
	switch dataPoint {
	case dataPoint:
		return FindCityLongitude
	default:
		return FindCityLongitude
	}
}

func GetLatudeForDataPoint(dataPoint string) LatFinder {
	switch dataPoint {
	case dataPoint:
		return FindCityLatitude
	default:
		return FindCityLatitude
	}
}
