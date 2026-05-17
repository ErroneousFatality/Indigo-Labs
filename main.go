package main

import (
	"IndigoLabs/Applications/Console"
	"IndigoLabs/Domain/Interfaces/DataSource"
	"IndigoLabs/Domain/Interfaces/DataStore"
	"IndigoLabs/Infrastructure/DataSource/Csv"
	"IndigoLabs/Infrastructure/DataStore/Memory"
)

func main() {
	var dataSource DataSource.IDataSource = &Csv.Reader{
		FilePath:   "C:\\Users\\Andre\\Desktop\\Potraga za poslom\\Applications\\Ljubljana\\Indigo Labs\\Interview project\\measurements.csv",
		Delimeter:  ';',
		DateFormat: "2006-01-02T15:04",
	}
	var dataStore DataStore.IDataStore = &Memory.Store{}
	Console.Startup(dataSource, dataStore)
}
