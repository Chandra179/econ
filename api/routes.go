package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter initializes the Gin router with API routes
// @title Stock Market API
// @version 1.0
// @description API for retrieving time series data from stock markets
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Configure CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// API versioning
	v1 := router.Group("/api/v1")
	{
		// Time series endpoints
		timeseries := v1.Group("/timeseries")
		{
			timeseries.GET("/:symbol", GetTimeSeriesForSymbol)
			timeseries.GET("/:symbol/:interval", GetTimeSeriesWithInterval)
		}
	}

	return router
}
