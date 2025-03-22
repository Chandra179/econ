package fundamental

import (
	"stock/common"
	"stock/config"
)

/*
	sector, industry, per, ebitda, eps
*/

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
	resp, err := common.MakeAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}

	// Convert generic response to CompanyOverviewResponse
	overview := &CompanyOverviewResponse{
		Symbol:                     getString(resp, "Symbol"),
		Name:                       getString(resp, "Name"),
		Description:                getString(resp, "Description"),
		Exchange:                   getString(resp, "Exchange"),
		Currency:                   getString(resp, "Currency"),
		Country:                    getString(resp, "Country"),
		Sector:                     getString(resp, "Sector"),
		Industry:                   getString(resp, "Industry"),
		Address:                    getString(resp, "Address"),
		FiscalYearEnd:              getString(resp, "FiscalYearEnd"),
		LatestQuarter:              getString(resp, "LatestQuarter"),
		MarketCapitalization:       getString(resp, "MarketCapitalization"),
		EBITDA:                     getString(resp, "EBITDA"),
		PERatio:                    getString(resp, "PERatio"),
		PEGRatio:                   getString(resp, "PEGRatio"),
		BookValue:                  getString(resp, "BookValue"),
		DividendPerShare:           getString(resp, "DividendPerShare"),
		DividendYield:              getString(resp, "DividendYield"),
		EPS:                        getString(resp, "EPS"),
		RevenuePerShareTTM:         getString(resp, "RevenuePerShareTTM"),
		ProfitMargin:               getString(resp, "ProfitMargin"),
		OperatingMarginTTM:         getString(resp, "OperatingMarginTTM"),
		ReturnOnAssetsTTM:          getString(resp, "ReturnOnAssetsTTM"),
		ReturnOnEquityTTM:          getString(resp, "ReturnOnEquityTTM"),
		RevenueTTM:                 getString(resp, "RevenueTTM"),
		GrossProfitTTM:             getString(resp, "GrossProfitTTM"),
		DilutedEPSTTM:              getString(resp, "DilutedEPSTTM"),
		QuarterlyEarningsGrowthYOY: getString(resp, "QuarterlyEarningsGrowthYOY"),
		QuarterlyRevenueGrowthYOY:  getString(resp, "QuarterlyRevenueGrowthYOY"),
		AnalystTargetPrice:         getString(resp, "AnalystTargetPrice"),
		TrailingPE:                 getString(resp, "TrailingPE"),
		ForwardPE:                  getString(resp, "ForwardPE"),
		PriceToSalesRatioTTM:       getString(resp, "PriceToSalesRatioTTM"),
		PriceToBookRatio:           getString(resp, "PriceToBookRatio"),
		EVToRevenue:                getString(resp, "EVToRevenue"),
		EVToEBITDA:                 getString(resp, "EVToEBITDA"),
		Beta:                       getString(resp, "Beta"),
		WeekHigh52:                 getString(resp, "52WeekHigh"),
		WeekLow52:                  getString(resp, "52WeekLow"),
		DayMovingAverage50:         getString(resp, "50DayMovingAverage"),
		DayMovingAverage200:        getString(resp, "200DayMovingAverage"),
		SharesOutstanding:          getString(resp, "SharesOutstanding"),
		SharesFloat:                getString(resp, "SharesFloat"),
		SharesShort:                getString(resp, "SharesShort"),
		SharesShortPriorMonth:      getString(resp, "SharesShortPriorMonth"),
		ShortRatio:                 getString(resp, "ShortRatio"),
		ShortPercentOutstanding:    getString(resp, "ShortPercentOutstanding"),
		ShortPercentFloat:          getString(resp, "ShortPercentFloat"),
		PercentInsiders:            getString(resp, "PercentInsiders"),
		PercentInstitutions:        getString(resp, "PercentInstitutions"),
		ForwardAnnualDividendRate:  getString(resp, "ForwardAnnualDividendRate"),
		ForwardAnnualDividendYield: getString(resp, "ForwardAnnualDividendYield"),
		PayoutRatio:                getString(resp, "PayoutRatio"),
		DividendDate:               getString(resp, "DividendDate"),
		ExDividendDate:             getString(resp, "ExDividendDate"),
		LastSplitFactor:            getString(resp, "LastSplitFactor"),
		LastSplitDate:              getString(resp, "LastSplitDate"),
	}

	return overview, nil
}

// Helper function to safely extract string values from the response map
func getString(data map[string]interface{}, key string) string {
	if val, ok := data[key]; ok {
		if strVal, ok := val.(string); ok {
			return strVal
		}
	}
	return ""
}
