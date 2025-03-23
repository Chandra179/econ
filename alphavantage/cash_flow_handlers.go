package alphavantage

import (
	"net/http"
	"stock/alphavantage/fundamental"
	"time"

	"github.com/gin-gonic/gin"
)

// CashFlowResponse defines the response format for cash flow data
// @Description Cash flow response data structure
type CashFlowResponse struct {
	Version   string                        `json:"version"`
	Timestamp string                        `json:"timestamp"`
	Symbol    string                        `json:"symbol"`
	Data      *fundamental.CashFlowResponse `json:"data"`
}

// GetCashFlow handles requests for cash flow data
// @Summary Get cash flow data for a specific symbol
// @Description Returns the cash flow data for the specified stock symbol
// @Tags fundamental
// @Produce json
// @Param symbol path string true "Stock symbol (e.g., AAPL, MSFT)"
// @Success 200 {object} CashFlowResponse "Successful operation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/fundamental/cash-flow/{symbol} [get]
func GetCashFlow(c *gin.Context) {
	symbol := c.Param("symbol")

	// Create params for the fundamental library
	params := fundamental.CashFlowParams{
		Symbol: symbol,
	}

	// Get cash flow data
	data, err := fundamental.GetCashFlow(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create response with versioning
	response := CashFlowResponse{
		Version:   "1.0", // TODO: Replace with config value once GetConfig() is implemented
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Symbol:    symbol,
		Data:      data,
	}

	c.JSON(http.StatusOK, response)
}
