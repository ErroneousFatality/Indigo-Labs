# Indigo Labs project assignment
A RESTFul API service for aggregating and displaying city temperature statistics.

## Configuration
A `config.json` file exists alongside the runtime executable file. Inside it you can change the settings for the source CSV file (its `FilePath`, `Delimeter` and `DateFormat`) and for the WebApi (its `Address`).

## Usage
Using the API endpoints, the user can accomplish the following:
 - **POST** `/Cities/ReloadData`: (Re)loads the data into the system from the source file.
 - **GET** `/Cities`: Displays the statistics for all the cities.
 - **GET** `/Cities/:name`: Displays the statistics a single city mathching the case-insensitive *name* parameter.
 - **GET** `/Cities/Averages?min=-5&max=25`: Displays the average temperature statistics for the cities whose average temperature is contained in the custom range `[min, max]`. Both *min* and *max* parameters are optional.

## Details
- The data is stored in memory and is lost upon the termination of the program.
- The data loading feature processes around 1.6 million CSV rows per second on a Ryzen 5 5600x (3.70GHz) processor with fast DDR4 RAM and NVME SSD storage.
- All other API requests execute in ~0s time for ~900 cities worth of records.