package stock

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
)

/*
	{
        "1. open": "324.7000",
        "2. high": "326.2000",
        "3. low": "320.6000",
        "4. close": "325.6000",
        "5. volume": "40252603"
	}
	---------------
	by API / CSV
	---------------
	interval (minutes)
	daily
	weekly, weekly adjusted
	monthly, monthly adjusted
*/

// TimeSeriesParams holds parameters for retrieving time series data
type TimeSeriesParams struct {
	Function      string // TIME_SERIES_INTRADAY, TIME_SERIES_DAILY, TIME_SERIES_WEEKLY, TIME_SERIES_MONTHLY
	Symbol        string
	Interval      string // Required for INTRADAY: 1min, 5min, 15min, 30min, 60min
	OutputSize    string // compact or full
	DataType      string // json or csv
	Adjusted      bool   // For adjusted time series
	ExtendedHours bool   // For intraday data
	Month         string // For intraday historical data in YYYY-MM format
}

// GetTimeSeries fetches time series data from Alpha Vantage API
func GetTimeSeries(params TimeSeriesParams) (map[string]interface{}, error) {
	baseURL := "https://www.alphavantage.co/query"

	// Get API key from environment
	apiKey := getAPIKeyFromEnv()

	// Building query parameters
	queryParams := map[string]string{
		"function": params.Function,
		"symbol":   params.Symbol,
		"apikey":   apiKey,
	}

	// Add optional parameters if provided
	if params.Interval != "" && params.Function == "TIME_SERIES_INTRADAY" {
		queryParams["interval"] = params.Interval
	}

	if params.OutputSize != "" {
		queryParams["outputsize"] = params.OutputSize
	}

	if params.DataType != "" {
		queryParams["datatype"] = params.DataType
	}

	if params.Function == "TIME_SERIES_INTRADAY" {
		if params.Month != "" {
			queryParams["month"] = params.Month
		}

		// Only include adjusted parameter if it's not the default value (true)
		if !params.Adjusted {
			queryParams["adjusted"] = "false"
		}

		if params.ExtendedHours {
			queryParams["extended_hours"] = "true"
		} else {
			queryParams["extended_hours"] = "false"
		}
	}

	// Make HTTP request and parse response
	resp, err := MakeAPIRequest(baseURL, queryParams)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// MakeAPIRequest performs the HTTP request to Alpha Vantage
func MakeAPIRequest(baseURL string, params map[string]string) (map[string]interface{}, error) {
	client := &http.Client{}

	// Build URL with query parameters
	url, err := buildRequestURL(baseURL, params)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse JSON response
	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}

// buildRequestURL builds the complete URL with query parameters
func buildRequestURL(baseURL string, params map[string]string) (string, error) {
	u, err := url.Parse(baseURL)
	if err != nil {
		return "", err
	}

	q := u.Query()
	for key, value := range params {
		q.Set(key, value)
	}

	u.RawQuery = q.Encode()
	return u.String(), nil
}

// getAPIKeyFromEnv gets the Alpha Vantage API key from environment
func getAPIKeyFromEnv() string {
	apiKey := os.Getenv("ALPHAVANTAGE_API_KEY")
	if apiKey == "" {
		// Fallback to "demo" if not set
		apiKey = "demo"
	}
	return apiKey
}

func getTimeSeries() {
	// Example usage:
	params := TimeSeriesParams{
		Function:   "TIME_SERIES_DAILY",
		Symbol:     "IBM",
		OutputSize: "compact",
		DataType:   "json",
	}

	data, err := GetTimeSeries(params)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Data retrieved successfully:", data)
}
