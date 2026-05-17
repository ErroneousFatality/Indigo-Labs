package DataStore

import "IndigoLabs/Domain/Entities/Cities"

type IDataStore interface {
	SetCities(map[string]*Cities.City)
	GetCities() map[string]*Cities.City
	GetCity(name string) (*Cities.City, bool)

	SetCityAverages([]Cities.CityAverage)
	GetCityAverages(min, max float32) []Cities.CityAverage
}
