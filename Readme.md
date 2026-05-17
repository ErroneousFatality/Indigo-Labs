# Indigo Labs project assignment
A RESTFul API service for aggregating and displaying city temperature statistics.

## Configuration
A `config.json` file exists alongside the runtime executable file. Inside it you can change the settings for the source CSV file (its `FilePath`, `Delimeter` and `DateFormat`) and for the WebApi (its `Address`).

## Usage
Using the API endpoints, the user can accomplish the following:
 - **POST** `/Cities/ReloadData`: (Re)loads the data into the system from the source file (takes about a second per one million CSV rows).
 - **GET** `/Cities`: Displays the statistics for all the cities.
 - **GET** `/Cities/:name`: Displays the statistics a single city mathching the case-insensitive *name* parameter.
 - **GET** `/Cities/Averages?min=-5&max=25`: Displays the average temperature statistics for the cities whose average temperature is contained in the custom range `[min, max]`. Both *min* and *max* parameters are optional.
