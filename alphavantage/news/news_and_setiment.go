package news

import (
	"encoding/json"
	"stock/common"
	"stock/config"
	"strconv"
)

type GetNewsAndSentimentParams struct {
	Tickers  string
	Topics   string
	TimeFrom string // YYYYMMDDTHHMM format
	TimeTo   string // YYYYMMDDTHHMM format
	Sort     string // LATEST, EARLIEST, or RELEVANCE
	Limit    int    // Default 50, max 1000
}

type Topic struct {
	Topic          string `json:"topic"`
	RelevanceScore string `json:"relevance_score"`
}

type FeedItem struct {
	Title                 string            `json:"title"`
	URL                   string            `json:"url"`
	TimePublished         string            `json:"time_published"`
	Authors               []string          `json:"authors"`
	Summary               string            `json:"summary"`
	Source                string            `json:"source"`
	CategoryWithin        string            `json:"category_within_source"`
	SourceDomain          string            `json:"source_domain"`
	Topics                []Topic           `json:"topics"`
	OverallSentimentScore float64           `json:"overall_sentiment_score"`
	OverallSentimentLabel string            `json:"overall_sentiment_label"`
	TickerSentiment       []TickerSentiment `json:"ticker_sentiment"`
}

type TickerSentiment struct {
	Ticker               string `json:"ticker"`
	RelevanceScore       string `json:"relevance_score"`
	TickerSentimentScore string `json:"ticker_sentiment_score"`
	TickerSentimentLabel string `json:"ticker_sentiment_label"`
}

type GetNewsAndSentimentResponse struct {
	Items                    []FeedItem `json:"feed"`
	ItemsCount               string     `json:"items"`
	SentimentScoreDefinition string     `json:"sentiment_score_definition"`
	RelevanceScoreDefinition string     `json:"relevance_score_definition"`
}

func GetNewsAndSentiment(params GetNewsAndSentimentParams) (*GetNewsAndSentimentResponse, error) {
	// Get API configuration
	cfg := config.GetConfig()

	// Building query parameters
	queryParams := map[string]string{
		"function": "NEWS_SENTIMENT",
		"apikey":   cfg.AlphaVantageAPIKey,
	}

	// Add optional parameters if they are provided
	if params.Tickers != "" {
		queryParams["tickers"] = params.Tickers
	}
	if params.Topics != "" {
		queryParams["topics"] = params.Topics
	}
	if params.TimeFrom != "" {
		queryParams["time_from"] = params.TimeFrom
	}
	if params.TimeTo != "" {
		queryParams["time_to"] = params.TimeTo
	}
	if params.Sort != "" {
		queryParams["sort"] = params.Sort
	}
	if params.Limit > 0 {
		queryParams["limit"] = strconv.Itoa(params.Limit)
	}

	// Make HTTP request
	respBody, err := common.GetAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	// Convert generic response to CompanyOverviewResponse
	resp := &GetNewsAndSentimentResponse{}
	if err := json.NewDecoder(respBody).Decode(resp); err != nil {
		return nil, err
	}

	return resp, nil
}
