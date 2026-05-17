package DataStore

type CityAverageFilter struct {
	Min *float32
	Max *float32
}

func NewCityAverageFilter(min, max float32) CityAverageFilter {
	return CityAverageFilter{
		Min: &min,
		Max: &max,
	}
}
