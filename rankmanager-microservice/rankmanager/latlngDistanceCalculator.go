package rankmanager

type latLngDistanceCalculator func(float32, float32, float32, float32) float32

func latlngDistCalculator(searchTermLat, searchTermLng, realTermLat, realTermLng float32) float32 {
	return (searchTermLat - realTermLat) + (searchTermLng - realTermLng)
}
