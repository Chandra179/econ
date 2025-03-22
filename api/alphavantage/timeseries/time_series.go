package timeseries

import (
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

// GetTimeSeries fetches time series data from Alpha Vantage API
func GetTimeSeries(params TimeSeriesParams) (map[string]interface{}, error) {
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
	resp, err := common.MakeAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
