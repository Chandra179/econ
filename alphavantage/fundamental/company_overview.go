package fundamental

import (
	"encoding/json"
	"stock/common"
	"stock/config"
)

// CompanyOverviewParams holds parameters for retrieving company overview data
type CompanyOverviewParams struct {
	Symbol string
}

// CompanyOverviewResponse defines the response format for company overview data
type CompanyOverviewResponse struct {
	Symbol                     string `json:"symbol"`
	Name                       string `json:"name"`
	Description                string `json:"description"`
	Exchange                   string `json:"exchange"`
	Currency                   string `json:"currency"`
	Country                    string `json:"country"`
	Sector                     string `json:"sector"`
	Industry                   string `json:"industry"`
	Address                    string `json:"address"`
	FiscalYearEnd              string `json:"fiscalYearEnd"`
	LatestQuarter              string `json:"latestQuarter"`
	MarketCapitalization       string `json:"marketCapitalization"`
	EBITDA                     string `json:"ebitda"`
	PERatio                    string `json:"peRatio"`
	PEGRatio                   string `json:"pegRatio"`
	BookValue                  string `json:"bookValue"`
	DividendPerShare           string `json:"dividendPerShare"`
	DividendYield              string `json:"dividendYield"`
	EPS                        string `json:"eps"`
	RevenuePerShareTTM         string `json:"revenuePerShareTTM"`
	ProfitMargin               string `json:"profitMargin"`
	OperatingMarginTTM         string `json:"operatingMarginTTM"`
	ReturnOnAssetsTTM          string `json:"returnOnAssetsTTM"`
	ReturnOnEquityTTM          string `json:"returnOnEquityTTM"`
	RevenueTTM                 string `json:"revenueTTM"`
	GrossProfitTTM             string `json:"grossProfitTTM"`
	DilutedEPSTTM              string `json:"dilutedEPSTTM"`
	QuarterlyEarningsGrowthYOY string `json:"quarterlyEarningsGrowthYOY"`
	QuarterlyRevenueGrowthYOY  string `json:"quarterlyRevenueGrowthYOY"`
	AnalystTargetPrice         string `json:"analystTargetPrice"`
	TrailingPE                 string `json:"trailingPE"`
	ForwardPE                  string `json:"forwardPE"`
	PriceToSalesRatioTTM       string `json:"priceToSalesRatioTTM"`
	PriceToBookRatio           string `json:"priceToBookRatio"`
	EVToRevenue                string `json:"evToRevenue"`
	EVToEBITDA                 string `json:"evToEBITDA"`
	Beta                       string `json:"beta"`
	WeekHigh52                 string `json:"52WeekHigh"`
	WeekLow52                  string `json:"52WeekLow"`
	DayMovingAverage50         string `json:"50DayMovingAverage"`
	DayMovingAverage200        string `json:"200DayMovingAverage"`
	SharesOutstanding          string `json:"sharesOutstanding"`
	SharesFloat                string `json:"sharesFloat"`
	SharesShort                string `json:"sharesShort"`
	SharesShortPriorMonth      string `json:"sharesShortPriorMonth"`
	ShortRatio                 string `json:"shortRatio"`
	ShortPercentOutstanding    string `json:"shortPercentOutstanding"`
	ShortPercentFloat          string `json:"shortPercentFloat"`
	PercentInsiders            string `json:"percentInsiders"`
	PercentInstitutions        string `json:"percentInstitutions"`
	ForwardAnnualDividendRate  string `json:"forwardAnnualDividendRate"`
	ForwardAnnualDividendYield string `json:"forwardAnnualDividendYield"`
	PayoutRatio                string `json:"payoutRatio"`
	DividendDate               string `json:"dividendDate"`
	ExDividendDate             string `json:"exDividendDate"`
	LastSplitFactor            string `json:"lastSplitFactor"`
	LastSplitDate              string `json:"lastSplitDate"`
}

// GetCompanyOverview fetches company overview data from Alpha Vantage API
func GetCompanyOverview(params CompanyOverviewParams) (*CompanyOverviewResponse, error) {
	// Get API configuration
	cfg := config.GetConfig()

	// Building query parameters
	queryParams := map[string]string{
		"function": "OVERVIEW",
		"symbol":   params.Symbol,
		"apikey":   cfg.AlphaVantageAPIKey,
	}

	// Make HTTP request
	respBody, err := common.GetAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	// Convert generic response to CompanyOverviewResponse
	overview := &CompanyOverviewResponse{
		Symbol: params.Symbol,
	}
	if err := json.NewDecoder(respBody).Decode(overview); err != nil {
		return nil, err
	}

	return overview, nil
}
