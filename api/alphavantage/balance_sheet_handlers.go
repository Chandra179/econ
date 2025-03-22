package alphavantage

import (
	"net/http"
	"stock/api/alphavantage/fundamental"
	"time"

	"github.com/gin-gonic/gin"
)

// BalanceSheetResponse defines the response format for balance sheet data
// @Description Balance sheet response data structure
type BalanceSheetResponse struct {
	Version   string                            `json:"version"`
	Timestamp string                            `json:"timestamp"`
	Symbol    string                            `json:"symbol"`
	Data      *fundamental.BalanceSheetResponse `json:"data"`
}

// GetBalanceSheet handles requests for balance sheet data
// @Summary Get balance sheet data for a specific symbol
// @Description Returns the balance sheet data for the specified stock symbol
// @Tags fundamental
// @Produce json
// @Param symbol path string true "Stock symbol (e.g., AAPL, MSFT)"
// @Success 200 {object} BalanceSheetResponse "Successful operation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/fundamental/balance-sheet/{symbol} [get]
func GetBalanceSheet(c *gin.Context) {
	symbol := c.Param("symbol")

	// Create params for the fundamental library
	params := fundamental.BalanceSheetParams{
		Symbol: symbol,
	}

	// Get balance sheet data
	data, err := fundamental.GetBalanceSheet(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create response with versioning
	response := BalanceSheetResponse{
		Version:   "1.0", // TODO: Replace with config value once GetConfig() is implemented
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Symbol:    symbol,
		Data:      data,
	}

	c.JSON(http.StatusOK, response)
}
