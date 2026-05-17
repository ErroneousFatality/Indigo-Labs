package CityUsecases

import "IndigoLabs/Domain/Entities/Cities"

type CityResponse struct {
	Name           string
	CelsiusMin     float32
	CelsiusMax     float32
	CelsiusAverage float32
}

func toResponse(city *Cities.City) CityResponse {
	return CityResponse{
		Name:           city.Name,
		CelsiusMin:     city.CelsiusMin,
		CelsiusMax:     city.CelsiusMax,
		CelsiusAverage: city.CelsiusAverage,
	}
}

func toResponsePointer(city *Cities.City) *CityResponse {
	response := toResponse(city)
	return &response
}
