package DataStore

import "IndigoLabs/Domain/Entities/Cities"

type IDataStore interface {
	SetCities([]Cities.City)

	GetCities() map[string]*Cities.City
	GetCity(name string) (*Cities.City, bool)
	GetCityAverages(filter CityAverageFilter) []Cities.CityAverage
}
