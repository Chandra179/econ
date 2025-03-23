package alphavantage

import (
	"net/http"
	"time"

	"stock/alphavantage/fundamental"

	"github.com/gin-gonic/gin"
)

// IncomeStatementResponse defines the response format for income statement data
// @Description Income statement response data structure
type IncomeStatementResponse struct {
	Version   string                               `json:"version"`
	Timestamp string                               `json:"timestamp"`
	Symbol    string                               `json:"symbol"`
	Data      *fundamental.IncomeStatementResponse `json:"data"`
}

// GetIncomeStatement handles requests for income statement data
// @Summary Get income statement data for a specific symbol
// @Description Returns the income statement data for the specified stock symbol
// @Tags fundamental
// @Produce json
// @Param symbol path string true "Stock symbol (e.g., AAPL, MSFT)"
// @Success 200 {object} IncomeStatementResponse "Successful operation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/fundamental/income-statement/{symbol} [get]
func GetIncomeStatement(c *gin.Context) {
	symbol := c.Param("symbol")

	// Create params for the fundamental library
	params := fundamental.IncomeStatementParams{
		Symbol: symbol,
	}

	// Get income statement data
	data, err := fundamental.GetIncomeStatement(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create response with versioning
	response := IncomeStatementResponse{
		Version:   "1.0", // TODO: Replace with config value once GetConfig() is implemented
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Symbol:    symbol,
		Data:      data,
	}

	c.JSON(http.StatusOK, response)
}
