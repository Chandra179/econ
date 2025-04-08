package timeseries

import (
	"encoding/json"
	"stock/common"
	"stock/config"
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

// TimeSeriesMetaData represents the Meta Data field in the API response
type TimeSeriesMetaData struct {
	Information   string `json:"1. Information"`
	Symbol        string `json:"2. Symbol"`
	LastRefreshed string `json:"3. Last Refreshed"`
	Interval      string `json:"4. Interval"`
	OutputSize    string `json:"5. Output Size"`
	TimeZone      string `json:"6. Time Zone"`
}

// TimeSeriesData represents the data point structure for each timestamp
type TimeSeriesData struct {
	Open   string `json:"1. open"`
	High   string `json:"2. high"`
	Low    string `json:"3. low"`
	Close  string `json:"4. close"`
	Volume string `json:"5. volume"`
}

// TimeSeriesResponse represents the complete API response structure
type TimeSeriesResponse struct {
	MetaData    TimeSeriesMetaData        `json:"Meta Data"`
	TimeSeries  map[string]TimeSeriesData `json:"Time Series (5min),omitempty"`
	DailyData   map[string]TimeSeriesData `json:"Time Series (Daily),omitempty"`
	WeeklyData  map[string]TimeSeriesData `json:"Weekly Time Series,omitempty"`
	MonthlyData map[string]TimeSeriesData `json:"Monthly Time Series,omitempty"`
}

// GetTimeSeries fetches time series data from Alpha Vantage API
func GetTimeSeries(params TimeSeriesParams) (*TimeSeriesResponse, error) {
	// Get API configuration
	cfg := config.GetConfig()

	// Building query parameters
	queryParams := map[string]string{
		"function": params.Function,
		"symbol":   params.Symbol,
		"apikey":   cfg.AlphaVantageAPIKey,
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
	respBody, err := common.GetAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	result := &TimeSeriesResponse{}
	if err := json.NewDecoder(respBody).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
