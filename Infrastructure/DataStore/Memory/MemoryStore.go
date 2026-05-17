package Memory

import (
	"IndigoLabs/Domain/Entities/Cities"
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
