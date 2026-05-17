package CityUsecases

import (
	"IndigoLabs/Domain/Entities/Cities"
	"IndigoLabs/Domain/Interfaces/DataSource"
	"IndigoLabs/Domain/Interfaces/DataStore"
	"log"
	"strings"
)

type CityUsecases struct {
	Source DataSource.IDataSource
	Store  DataStore.IDataStore
}

func (this *CityUsecases) RecreateData() {
	aggregates := make(map[string]*cityAggregate)

	for result := range this.Source.ReadStream() {
		if result.Err != nil {
			log.Println(result.Err)
			continue
		}
		measurement := result.Data

		aggregate, ok := aggregates[measurement.City]
		if ok {
			aggregate.AddMeasurement(measurement.CelsiusAverage)
		} else {
			aggregates[measurement.City] = newCityAggregate(measurement.CelsiusAverage)
		}
	}

	cities := make([]Cities.City, len(aggregates))
	i := 0
	for name, aggregate := range aggregates {
		cities[i] = Cities.City{
			Name:           name,
			CelsiusMin:     aggregate.CelsiusMin,
			CelsiusMax:     aggregate.CelsiusMax,
			CelsiusAverage: aggregate.GetCelsiusAverage(),
		}
		i++
	}
	this.Store.SetCities(cities)
}

func (this *CityUsecases) GetAllCities() []CityResponse {
	cities := this.Store.GetCities()

	responses := make([]CityResponse, len(cities))
	i := 0
	for _, city := range cities {
		responses[i] = toResponse(city)
		i++
	}

	return responses
}

func (this *CityUsecases) GetCity(name string) *CityResponse {
	keyname := strings.ToUpper(name)
	city, ok := this.Store.GetCity(keyname)
	if !ok {
		return nil
	}
	return toResponsePointer(city)
}

func (this *CityUsecases) GetCityAverages(min float32, max float32) []Cities.CityAverage {
	return this.Store.GetCityAverages(min, max)
}
