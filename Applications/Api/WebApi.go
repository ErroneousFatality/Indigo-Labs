package Api

import (
	"IndigoLabs/Domain/Interfaces/DataSource"
	"IndigoLabs/Domain/Interfaces/DataStore"
	CityUsecases "IndigoLabs/Domain/Usecases/Cities"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Startup(dataSource DataSource.IDataSource, dataStore DataStore.IDataStore, address string) {
	// Usecases
	cityUsecases := &CityUsecases.CityUsecases{
		Source: dataSource,
		Store:  dataStore,
	}

	// Api
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(CORSMiddleware())

	router.POST("/Cities/ReloadData", func(context *gin.Context) {
		cityUsecases.RecreateData()
		context.Status(http.StatusOK)
	})
	router.GET("/Cities", func(context *gin.Context) {
		cities := cityUsecases.GetAllCities()
		context.JSON(http.StatusOK, cities)
	})
	router.GET("/Cities/:name", func(context *gin.Context) {
		cityName := context.Param("name")
		city := cityUsecases.GetCity(cityName)
		if city == nil {
			context.JSON(http.StatusNotFound, gin.H{"error": "City not found"})
		} else {
			context.JSON(http.StatusOK, city)
		}
	})
	router.GET("/Cities/Averages", func(context *gin.Context) {
		var filter DataStore.CityAverageFilter
		if err := context.ShouldBindQuery(&filter); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{
				"error":   "Invalid query parameters",
				"details": err.Error(),
			})
			return
		}
		cityAverages := cityUsecases.GetCityAverages(filter)
		context.JSON(http.StatusOK, cityAverages)
	})
	router.Run(address)
}
