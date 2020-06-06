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
