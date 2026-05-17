package DataSource

import "time"

type DailyCityTemperature struct {
	Date           time.Time `csv:"datetime"`
	City           string    `csv:"city"`
	CelsiusAverage float32   `csv:"temp_celsius"`
}
