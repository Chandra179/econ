package api

import (
	"log"
	"os"
	"os/exec"
	"stock/api/alphavantage"

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
			timeseries.GET("/:symbol", alphavantage.GetTimeSeriesForSymbol)
			timeseries.GET("/:symbol/:interval", alphavantage.GetTimeSeriesWithInterval)
		}

		// Fundamental data endpoints
		fundamental := v1.Group("/fundamental")
		{
			fundamental.GET("/balance-sheet/:symbol", alphavantage.GetBalanceSheet)
			fundamental.GET("/cash-flow/:symbol", alphavantage.GetCashFlow)
			fundamental.GET("/income-statement/:symbol", alphavantage.GetIncomeStatement)
			fundamental.GET("/company-overview/:symbol", alphavantage.GetCompanyOverview)
		}
	}

	return router
}

func init() {
	// Only in development mode
	if os.Getenv("GO_ENV") == "development" {
		cmd := exec.Command("swag", "init", "-g", "api/routes.go", "-o", "docs")
		err := cmd.Run()
		if err != nil {
			log.Println("Warning: Swagger docs generation failed:", err)
		}
	}
}
