package CityUsecases

type cityAggregate struct {
	MeasurementsCount uint32
	CelsiusMax        float32
	CelsiusMin        float32
	CelsiusSum        float64
	CelsiusAverage    float32
}

func newCityAggregate(celsius float32) *cityAggregate {
	return &cityAggregate{
		MeasurementsCount: 1,
		CelsiusMax:        celsius,
		CelsiusMin:        celsius,
		CelsiusSum:        float64(celsius),
		CelsiusAverage:    celsius,
	}
}

func (stats *cityAggregate) AddMeasurement(celsius float32) {
	stats.MeasurementsCount++
	if celsius > stats.CelsiusMax {
		stats.CelsiusMax = celsius
	}
	if celsius < stats.CelsiusMin {
		stats.CelsiusMin = celsius
	}
	stats.CelsiusSum += float64(celsius)
	stats.CelsiusAverage = float32(stats.CelsiusSum / float64(stats.MeasurementsCount))
}
