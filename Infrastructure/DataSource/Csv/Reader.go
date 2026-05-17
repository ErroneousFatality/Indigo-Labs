package Csv

import (
	"IndigoLabs/Domain/Interfaces/DataSource"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type Reader struct {
	FilePath   string
	Delimeter  rune
	DateFormat string
}

func (reader *Reader) ReadStream() <-chan DataSource.StreamResult {
	return StreamFile(reader.FilePath, reader.Delimeter, reader.DateFormat)
}

func StreamFile(filePath string, delimeter rune, dateFormat string) <-chan DataSource.StreamResult {
	outflow := make(chan DataSource.StreamResult, 10)
	fail := func(message string, err error) {
		outflow <- DataSource.StreamResult{Err: fmt.Errorf("%s: %w", message, err)}
	}
	go func() {
		defer close(outflow)

		// Open file
		file, err := os.Open(filePath)
		if err != nil {
			fail("Failed to open file", err)
			return
		}
		defer file.Close()

		// Create CSV reader
		csvReader := csv.NewReader(file)
		csvReader.Comma = delimeter

		// Skip the CSV header
		_, err = csvReader.Read()
		if err != nil {
			if err != io.EOF {
				fail("Failed to read CSV row #1", err)
			}
			return
		}

		// Parse and stream rows
		var rowNumber uint32 = 2
		for ; ; rowNumber++ {
			var row []string
			row, err := csvReader.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				fail(fmt.Sprintf("Failed to read CSV row #%d", rowNumber), err)
				continue
			}

			var data DataSource.DailyCityTemperature
			data, err = parseRow(row, dateFormat)
			if err != nil {
				fail(fmt.Sprintf("Failed to parse CSV row #%d", rowNumber), err)
				continue
			}

			outflow <- DataSource.StreamResult{Data: data}
		}
	}()
	return outflow
}

func parseRow(row []string, dateFormat string) (DataSource.DailyCityTemperature, error) {
	var date time.Time
	date, err := time.Parse(dateFormat, row[0])
	if err != nil {
		return DataSource.DailyCityTemperature{}, fmt.Errorf("Failed to parse date: %w", err)
	}

	var celsiusAverage float32
	celsiusAverage64, err := strconv.ParseFloat(row[2], 32)
	if err != nil {
		return DataSource.DailyCityTemperature{}, fmt.Errorf("Failed to parse celsius average: %w", err)
	}
	celsiusAverage = float32(celsiusAverage64)

	var data = DataSource.DailyCityTemperature{
		Date:           date,
		City:           row[1],
		CelsiusAverage: celsiusAverage,
	}
	return data, nil
}
