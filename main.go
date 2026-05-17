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
	for result := range reader.ReadStream() {
		if result.Err != nil {
			log.Println(result.Err)
			continue
		}
		log.Println(result.Data)
	var dataStore DataStore.IDataStore = &Memory.Store{}
	}
}
