package alphavantage

import (
	"net/http"
	"stock/api/alphavantage/fundamental"
	"time"

	"github.com/gin-gonic/gin"
)

// CompanyOverviewResponse defines the response format for company overview data
// @Description Company overview response data structure
type CompanyOverviewResponse struct {
	Version   string                               `json:"version"`
	Timestamp string                               `json:"timestamp"`
	Symbol    string                               `json:"symbol"`
	Data      *fundamental.CompanyOverviewResponse `json:"data"`
}

// GetCompanyOverview handles requests for company overview data
// @Summary Get company overview data for a specific symbol
// @Description Returns the company overview data for the specified stock symbol (sector, industry, PE ratio, EBITDA, and more)
// @Tags fundamental
// @Produce json
// @Param symbol path string true "Stock symbol (e.g., AAPL, MSFT)"
// @Success 200 {object} CompanyOverviewResponse "Successful operation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/fundamental/company-overview/{symbol} [get]
func GetCompanyOverview(c *gin.Context) {
	symbol := c.Param("symbol")

	// Create params for the fundamental library
	params := fundamental.CompanyOverviewParams{
		Symbol: symbol,
	}

	// Get company overview data
	data, err := fundamental.GetCompanyOverview(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create response with versioning
	response := CompanyOverviewResponse{
		Version:   "1.0", // TODO: Replace with config value once GetConfig() is implemented
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Symbol:    symbol,
		Data:      data,
	}

	c.JSON(http.StatusOK, response)
}
