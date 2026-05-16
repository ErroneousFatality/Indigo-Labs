package main

import (
	"IndigoLabs/DataSource"
	"IndigoLabs/DataSource/Implementations/Csv"
	"log"
)

func main() {
	var reader DataSource.IDataSource = Csv.CsvReader{
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
	}
}
