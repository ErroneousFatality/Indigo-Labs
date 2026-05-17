package main

import (
	"IndigoLabs/Applications/Api"
	"IndigoLabs/Domain/Interfaces/DataSource"
	"IndigoLabs/Domain/Interfaces/DataStore"
	"IndigoLabs/Infrastructure/DataSource/Csv"
	"IndigoLabs/Infrastructure/DataStore/Memory"
	"fmt"
	"unicode/utf8"

	"github.com/spf13/viper"
)

func main() {
	// Config
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("Failed to read the config file: %w", err))
	}

	// Data source
	csvDelimeter, _ := utf8.DecodeRuneInString(viper.GetString("DataSource.Csv.Delimeter"))
	var dataSource DataSource.IDataSource = &Csv.Reader{
		FilePath:   viper.GetString("DataSource.Csv.FilePath"),
		Delimeter:  csvDelimeter,
		DateFormat: viper.GetString("DataSource.Csv.DateFormat"),
	}

	// Data store
	var dataStore DataStore.IDataStore = &Memory.Store{}

	// Web API
	webApiAddress := viper.GetString("WebApi.Address")
	Api.Startup(dataSource, dataStore, webApiAddress)
}
