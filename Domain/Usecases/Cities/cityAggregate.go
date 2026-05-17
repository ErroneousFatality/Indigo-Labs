package CityUsecases

type cityAggregate struct {
	MeasurementsCount uint32
	CelsiusMax        float32
	CelsiusMin        float32
	CelsiusSum        float64
}

func newCityAggregate(celsius float32) *cityAggregate {
	return &cityAggregate{
		MeasurementsCount: 1,
		CelsiusMax:        celsius,
		CelsiusMin:        celsius,
		CelsiusSum:        float64(celsius),
	}
}

func (city *cityAggregate) AddMeasurement(celsius float32) {
	city.MeasurementsCount++
	if celsius > city.CelsiusMax {
		city.CelsiusMax = celsius
	}
	if celsius < city.CelsiusMin {
		city.CelsiusMin = celsius
	}
	city.CelsiusSum += float64(celsius)
}

func (city *cityAggregate) GetCelsiusAverage() float32 {
	return float32(city.CelsiusSum / float64(city.MeasurementsCount))
}
