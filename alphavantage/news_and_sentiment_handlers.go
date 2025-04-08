package alphavantage

import (
	"net/http"
	"stock/alphavantage/news"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// NewsAndSentimentResponse defines the response format for news and sentiment data
// @Description News and sentiment response data structure
type NewsAndSentimentResponse struct {
	Version   string                            `json:"version"`
	Timestamp string                            `json:"timestamp"`
	Data      *news.GetNewsAndSentimentResponse `json:"data"`
}

// GetNewsAndSentiment handles requests for news and sentiment data
// @Summary Get news and sentiment data for specified parameters
// @Description Returns news articles and sentiment analysis based on tickers, topics, and time range
// @Tags news
// @Produce json
// @Param tickers query string false "Comma-separated list of stock symbols (e.g., AAPL,MSFT)"
// @Param topics query string false "Comma-separated list of topics"
// @Param time_from query string false "Start time in YYYYMMDDTHHMM format"
// @Param time_to query string false "End time in YYYYMMDDTHHMM format"
// @Param sort query string false "Sort order: LATEST, EARLIEST, or RELEVANCE"
// @Param limit query int false "Number of results (default: 50, max: 1000)"
// @Success 200 {object} NewsAndSentimentResponse "Successful operation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /v1/news/sentiment [get]
func GetNewsAndSentiment(c *gin.Context) {
	// Extract query parameters
	params := news.GetNewsAndSentimentParams{
		Tickers:  c.Query("tickers"),
		Topics:   c.Query("topics"),
		TimeFrom: c.Query("time_from"),
		TimeTo:   c.Query("time_to"),
		Sort:     c.Query("sort"),
	}

	// Parse limit if provided
	if limitStr := c.Query("limit"); limitStr != "" {
		limit, err := strconv.Atoi(limitStr)
		if err == nil {
			params.Limit = limit
		}
	}

	// Get news and sentiment data
	data, err := news.GetNewsAndSentiment(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	// Create response with versioning
	response := NewsAndSentimentResponse{
		Version:   "1.0", // TODO: Replace with config value once GetConfig() is implemented
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		Data:      data,
	}

	c.JSON(http.StatusOK, response)
}
