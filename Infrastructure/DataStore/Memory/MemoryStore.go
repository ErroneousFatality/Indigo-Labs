package Memory

import (
	"IndigoLabs/Domain/Entities/Cities"
	"cmp"
	"slices"
)

type Store struct {
	cities       map[string]*Cities.City
	cityAverages []Cities.CityAverage
}

func (store *Store) SetCities(cities map[string]*Cities.City) {
	store.cities = cities
}

func (store *Store) SetCityAverages(cityAverages []Cities.CityAverage) {
	slices.SortFunc(cityAverages, func(a, b Cities.CityAverage) int {
		return cmp.Compare(a.CelsiusAverage, b.CelsiusAverage)
	})
	store.cityAverages = cityAverages
}

func (store *Store) GetCities() map[string]*Cities.City {
	return store.cities
}

func (store *Store) GetCity(name string) (*Cities.City, bool) {
	city, ok := store.cities[name]
	return city, ok
}

func (store *Store) GetCityAverages(min, max float32) []Cities.CityAverage {
	start, _ := slices.BinarySearchFunc(store.cityAverages, min, func(city Cities.CityAverage, celsiusAverage float32) int {
		return cmp.Compare(city.CelsiusAverage, celsiusAverage)
	})
	end, _ := slices.BinarySearchFunc(store.cityAverages, max, func(city Cities.CityAverage, celsiusAverage float32) int {
		return cmp.Compare(city.CelsiusAverage, celsiusAverage)
	})
	result := store.cityAverages[start:end]
	return result
}
