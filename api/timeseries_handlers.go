package api

import (
	"net/http"
	"strconv"
	"time"

	"stock/stock"

	"github.com/gin-gonic/gin"
)

// TimeSeriesParams holds parameters for retrieving time series data
type TimeSeriesParams struct {
	Function      string // TIME_SERIES_INTRADAY, TIME_SERIES_DAILY, TIME_SERIES_WEEKLY, TIME_SERIES_MONTHLY
	Symbol        string
	Interval      string // Required for INTRADAY: 1min, 5min, 15min, 30min, 60min
	OutputSize    string // compact or full
	DataType      string // json or csv
	ExtendedHours bool   // For intraday data
	Adjusted      bool   // For intraday data, default is true
	Month         string // For intraday historical data in YYYY-MM format
}

// TimeSeriesResponse defines the response format for time series data
// @Description Time series response data structure
type TimeSeriesResponse struct {
	Version   string                 `json:"version"`
	Timestamp string                 `json:"timestamp"`
	Symbol    string                 `json:"symbol"`
	Interval  string                 `json:"interval,omitempty"`
	Data      map[string]interface{} `json:"data"`
}

// GetTimeSeriesForSymbol handles requests for time series data for a specific symbol
// @Summary Get time series data for a specific symbol
// @Description Returns time series data for the specified stock symbol
// @Tags timeseries
// @Produce json
// @Param symbol path string true "Stock symbol (e.g., AAPL, MSFT)"
// @Param function query string false "Time series function" Enums(TIME_SERIES_DAILY, TIME_SERIES_WEEKLY, TIME_SERIES_MONTHLY) default(TIME_SERIES_DAILY)
// @Param outputsize query string false "Amount of data to return" Enums(compact, full) default(compact)
// @Param datatype query string false "Data type for response" Enums(json, csv) default(json)
// @Success 200 {object} TimeSeriesResponse "Successful operation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/timeseries/{symbol} [get]
func GetTimeSeriesForSymbol(c *gin.Context) {
	symbol := c.Param("symbol")
	function := c.DefaultQuery("function", "TIME_SERIES_DAILY")
	outputSize := c.DefaultQuery("outputsize", "compact")
	dataType := c.DefaultQuery("datatype", "json")

	// Create params for the stock library
	params := TimeSeriesParams{
		Function:   function,
		Symbol:     symbol,
		OutputSize: outputSize,
		DataType:   dataType,
	}

	// Get time series data
	data, err := getTimeSeriesData(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create response with versioning
	response := TimeSeriesResponse{
		Version:   GetConfig().DefaultVersion,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Symbol:    symbol,
		Data:      data,
	}

	c.JSON(http.StatusOK, response)
}

// GetTimeSeriesWithInterval handles requests for time series data with interval
// @Summary Get intraday time series data with specific interval
// @Description Returns intraday time series data with the specified interval
// @Tags timeseries
// @Produce json
// @Param symbol path string true "Stock symbol (e.g., AAPL, MSFT)"
// @Param interval path string true "Time interval for data" Enums(1min, 5min, 15min, 30min, 60min)
// @Param outputsize query string false "Amount of data to return" Enums(compact, full) default(compact)
// @Param datatype query string false "Data type for response" Enums(json, csv) default(json)
// @Param extended_hours query boolean false "Whether to include extended hours data" default(false)
// @Param adjusted query boolean false "Whether to adjust for split and dividend events" default(true)
// @Param month query string false "Month for historical intraday data (YYYY-MM format)"
// @Success 200 {object} TimeSeriesResponse "Successful operation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/timeseries/{symbol}/{interval} [get]
func GetTimeSeriesWithInterval(c *gin.Context) {
	symbol := c.Param("symbol")
	interval := c.Param("interval")
	function := "TIME_SERIES_INTRADAY" // Forced for interval-based queries
	outputSize := c.DefaultQuery("outputsize", "compact")
	dataType := c.DefaultQuery("datatype", "json")

	// Create params for the stock library
	params := TimeSeriesParams{
		Function:   function,
		Symbol:     symbol,
		Interval:   interval,
		OutputSize: outputSize,
		DataType:   dataType,
	}

	if extendedHours, err := strconv.ParseBool(c.DefaultQuery("extended_hours", "false")); err == nil {
		params.ExtendedHours = extendedHours
	}

	// Set adjusted parameter, default is true
	adjusted := true
	if adjustedParam, exists := c.GetQuery("adjusted"); exists {
		if parsedAdjusted, err := strconv.ParseBool(adjustedParam); err == nil {
			adjusted = parsedAdjusted
		}
	}
	params.Adjusted = adjusted

	// Optional month parameter for historical intraday data
	month := c.DefaultQuery("month", "")
	if month != "" {
		params.Month = month
	}

	// Get time series data
	data, err := getTimeSeriesData(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create response with versioning
	response := TimeSeriesResponse{
		Version:   GetConfig().DefaultVersion,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Symbol:    symbol,
		Interval:  interval,
		Data:      data,
	}

	c.JSON(http.StatusOK, response)
}

// getTimeSeriesData fetches time series data from Alpha Vantage API
func getTimeSeriesData(params TimeSeriesParams) (map[string]interface{}, error) {
	// Convert API params to stock package params
	stockParams := stock.TimeSeriesParams{
		Function:      params.Function,
		Symbol:        params.Symbol,
		Interval:      params.Interval,
		OutputSize:    params.OutputSize,
		DataType:      params.DataType,
		ExtendedHours: params.ExtendedHours,
		Month:         params.Month,
	}

	// Only set Adjusted if it's not the default value (true)
	// The stock.GetTimeSeries function will only include the adjusted parameter
	// in the API call if it's explicitly set to false
	stockParams.Adjusted = params.Adjusted

	// Use the library function directly
	return stock.GetTimeSeries(stockParams)
}
