package Memory

import (
	"IndigoLabs/Domain/Entities/Cities"
	"IndigoLabs/Domain/Interfaces/DataStore"
	"cmp"
	"slices"
	"strings"
)

type Store struct {
	cities       map[string]*Cities.City
	cityAverages []Cities.CityAverage
}

func (store *Store) SetCities(cities []Cities.City) {
	store.cities = make(map[string]*Cities.City, len(cities))
	store.cityAverages = make([]Cities.CityAverage, len(cities))
	for index, city := range cities {
		keyname := strings.ToUpper(city.Name)
		store.cities[keyname] = &city
		store.cityAverages[index] = Cities.CityAverage{
			Name:           city.Name,
			CelsiusAverage: city.CelsiusAverage,
		}
	}
	slices.SortFunc(store.cityAverages, func(a, b Cities.CityAverage) int {
		return cmp.Compare(a.CelsiusAverage, b.CelsiusAverage)
	})
}

func (store *Store) GetCities() map[string]*Cities.City {
	return store.cities
}

func (store *Store) GetCity(name string) (*Cities.City, bool) {
	city, ok := store.cities[name]
	return city, ok
}

func (store *Store) GetCityAverages(filter DataStore.CityAverageFilter) []Cities.CityAverage {
	var start int
	if filter.Min == nil {
		start = 0
	} else {
		start, _ = slices.BinarySearchFunc(store.cityAverages, *filter.Min, func(city Cities.CityAverage, celsiusAverageMin float32) int {
			return cmp.Compare(city.CelsiusAverage, celsiusAverageMin)
		})
	}

	var end int
	if filter.Max == nil {
		end = len(store.cityAverages)
	} else {
		end, _ = slices.BinarySearchFunc(store.cityAverages, *filter.Max, func(city Cities.CityAverage, celsiusAverageMax float32) int {
			return cmp.Compare(city.CelsiusAverage, celsiusAverageMax)
		})
	}

	result := store.cityAverages[start : end+1]
	return result
}
