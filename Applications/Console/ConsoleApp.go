package Console

import (
	"IndigoLabs/Domain/Interfaces/DataSource"
	"IndigoLabs/Domain/Interfaces/DataStore"
	CityUsecases "IndigoLabs/Domain/Usecases/Cities"
	"log"
)

func Startup(dataSource DataSource.IDataSource, dataStore DataStore.IDataStore) {
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
