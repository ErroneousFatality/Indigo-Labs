package main

import (
	"IndigoLabs/Domain/Interfaces/DataSource"
	"IndigoLabs/Domain/Interfaces/DataStore"
	CityUsecases "IndigoLabs/Domain/Usecases/Cities"
	"IndigoLabs/Infrastructure/DataSource/Csv"
	"IndigoLabs/Infrastructure/DataStore/Memory"
	"log"
)

func main() {
	var dataSource DataSource.IDataSource = &Csv.Reader{
		FilePath:   "C:\\Users\\Andre\\Desktop\\Potraga za poslom\\Applications\\Ljubljana\\Indigo Labs\\Interview project\\measurements.csv",
		Delimeter:  ';',
		DateFormat: "2006-01-02T15:04",
	}
	var dataStore DataStore.IDataStore = &Memory.Store{}
	cityUsecases := &CityUsecases.CityUsecases{
		Source: dataSource,
		Store:  dataStore,
	}

	cityUsecases.RecreateData()

	cities := cityUsecases.GetAllCities()
	log.Println(len(cities))

	city := cityUsecases.GetCity("ljuBljAna")
	log.Println(city)

	cityAverages := cityUsecases.GetCityAverages(0, 25)
	log.Println(cityAverages)
}
