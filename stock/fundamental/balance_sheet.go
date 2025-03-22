package fundamental

import (
	"stock/common"
	"stock/config"
)

// BalanceSheetParams holds parameters for retrieving balance sheet data
type BalanceSheetParams struct {
	Symbol string // Required: Stock symbol (e.g., IBM)
}

// BalanceSheetResponse defines the response format for balance sheet data
type BalanceSheetResponse struct {
	Symbol           string               `json:"symbol"`
	AnnualReports    []BalanceSheetReport `json:"annualReports"`
	QuarterlyReports []BalanceSheetReport `json:"quarterlyReports"`
}

// BalanceSheetReport represents a single balance sheet report
type BalanceSheetReport struct {
	FiscalDateEnding                       string `json:"fiscalDateEnding"`
	ReportedCurrency                       string `json:"reportedCurrency"`
	TotalAssets                            string `json:"totalAssets"`
	TotalCurrentAssets                     string `json:"totalCurrentAssets"`
	CashAndCashEquivalentsAtCarryingValue  string `json:"cashAndCashEquivalentsAtCarryingValue"`
	CashAndShortTermInvestments            string `json:"cashAndShortTermInvestments"`
	Inventory                              string `json:"inventory"`
	CurrentNetReceivables                  string `json:"currentNetReceivables"`
	TotalNonCurrentAssets                  string `json:"totalNonCurrentAssets"`
	PropertyPlantEquipment                 string `json:"propertyPlantEquipment"`
	AccumulatedDepreciationAmortizationPPE string `json:"accumulatedDepreciationAmortizationPPE"`
	IntangibleAssets                       string `json:"intangibleAssets"`
	IntangibleAssetsExcludingGoodwill      string `json:"intangibleAssetsExcludingGoodwill"`
	Goodwill                               string `json:"goodwill"`
	Investments                            string `json:"investments"`
	LongTermInvestments                    string `json:"longTermInvestments"`
	ShortTermInvestments                   string `json:"shortTermInvestments"`
	OtherCurrentAssets                     string `json:"otherCurrentAssets"`
	OtherNonCurrentAssets                  string `json:"otherNonCurrentAssets"`
	TotalLiabilities                       string `json:"totalLiabilities"`
	TotalCurrentLiabilities                string `json:"totalCurrentLiabilities"`
	CurrentAccountsPayable                 string `json:"currentAccountsPayable"`
	DeferredRevenue                        string `json:"deferredRevenue"`
	CurrentDebt                            string `json:"currentDebt"`
	ShortTermDebt                          string `json:"shortTermDebt"`
	TotalNonCurrentLiabilities             string `json:"totalNonCurrentLiabilities"`
	CapitalLeaseObligations                string `json:"capitalLeaseObligations"`
	LongTermDebt                           string `json:"longTermDebt"`
	CurrentLongTermDebt                    string `json:"currentLongTermDebt"`
	LongTermDebtNoncurrent                 string `json:"longTermDebtNoncurrent"`
	ShortLongTermDebtTotal                 string `json:"shortLongTermDebtTotal"`
	OtherCurrentLiabilities                string `json:"otherCurrentLiabilities"`
	OtherNonCurrentLiabilities             string `json:"otherNonCurrentLiabilities"`
	TotalShareholderEquity                 string `json:"totalShareholderEquity"`
	TreasuryStock                          string `json:"treasuryStock"`
	RetainedEarnings                       string `json:"retainedEarnings"`
	CommonStock                            string `json:"commonStock"`
	CommonStockSharesOutstanding           string `json:"commonStockSharesOutstanding"`
}

// GetBalanceSheet fetches balance sheet data from Alpha Vantage API
func GetBalanceSheet(params BalanceSheetParams) (*BalanceSheetResponse, error) {
	// Get API configuration
	cfg := config.GetConfig()

	// Building query parameters
	queryParams := map[string]string{
		"function": "BALANCE_SHEET",
		"symbol":   params.Symbol,
		"apikey":   cfg.AlphaVantageAPIKey,
	}

	// Make HTTP request
	resp, err := common.MakeAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}

	// Convert generic response to BalanceSheetResponse
	balanceSheetResp := &BalanceSheetResponse{
		Symbol: params.Symbol,
	}

	// Extract annual reports
	if annualReports, ok := resp["annualReports"].([]interface{}); ok {
		for _, report := range annualReports {
			if reportMap, ok := report.(map[string]interface{}); ok {
				balanceSheetReport := mapToBalanceSheetReport(reportMap)
				balanceSheetResp.AnnualReports = append(balanceSheetResp.AnnualReports, balanceSheetReport)
			}
		}
	}

	// Extract quarterly reports
	if quarterlyReports, ok := resp["quarterlyReports"].([]interface{}); ok {
		for _, report := range quarterlyReports {
			if reportMap, ok := report.(map[string]interface{}); ok {
				balanceSheetReport := mapToBalanceSheetReport(reportMap)
				balanceSheetResp.QuarterlyReports = append(balanceSheetResp.QuarterlyReports, balanceSheetReport)
			}
		}
	}

	return balanceSheetResp, nil
}

// mapToBalanceSheetReport converts a map to a BalanceSheetReport struct
func mapToBalanceSheetReport(reportMap map[string]interface{}) BalanceSheetReport {
	report := BalanceSheetReport{}

	// Helper function to safely extract string values
	getString := func(m map[string]interface{}, key string) string {
		if val, ok := m[key]; ok {
			if strVal, ok := val.(string); ok {
				return strVal
			}
		}
		return ""
	}

	// Extract all fields
	report.FiscalDateEnding = getString(reportMap, "fiscalDateEnding")
	report.ReportedCurrency = getString(reportMap, "reportedCurrency")
	report.TotalAssets = getString(reportMap, "totalAssets")
	report.TotalCurrentAssets = getString(reportMap, "totalCurrentAssets")
	report.CashAndCashEquivalentsAtCarryingValue = getString(reportMap, "cashAndCashEquivalentsAtCarryingValue")
	report.CashAndShortTermInvestments = getString(reportMap, "cashAndShortTermInvestments")
	report.Inventory = getString(reportMap, "inventory")
	report.CurrentNetReceivables = getString(reportMap, "currentNetReceivables")
	report.TotalNonCurrentAssets = getString(reportMap, "totalNonCurrentAssets")
	report.PropertyPlantEquipment = getString(reportMap, "propertyPlantEquipment")
	report.AccumulatedDepreciationAmortizationPPE = getString(reportMap, "accumulatedDepreciationAmortizationPPE")
	report.IntangibleAssets = getString(reportMap, "intangibleAssets")
	report.IntangibleAssetsExcludingGoodwill = getString(reportMap, "intangibleAssetsExcludingGoodwill")
	report.Goodwill = getString(reportMap, "goodwill")
	report.Investments = getString(reportMap, "investments")
	report.LongTermInvestments = getString(reportMap, "longTermInvestments")
	report.ShortTermInvestments = getString(reportMap, "shortTermInvestments")
	report.OtherCurrentAssets = getString(reportMap, "otherCurrentAssets")
	report.OtherNonCurrentAssets = getString(reportMap, "otherNonCurrentAssets")
	report.TotalLiabilities = getString(reportMap, "totalLiabilities")
	report.TotalCurrentLiabilities = getString(reportMap, "totalCurrentLiabilities")
	report.CurrentAccountsPayable = getString(reportMap, "currentAccountsPayable")
	report.DeferredRevenue = getString(reportMap, "deferredRevenue")
	report.CurrentDebt = getString(reportMap, "currentDebt")
	report.ShortTermDebt = getString(reportMap, "shortTermDebt")
	report.TotalNonCurrentLiabilities = getString(reportMap, "totalNonCurrentLiabilities")
	report.CapitalLeaseObligations = getString(reportMap, "capitalLeaseObligations")
	report.LongTermDebt = getString(reportMap, "longTermDebt")
	report.CurrentLongTermDebt = getString(reportMap, "currentLongTermDebt")
	report.LongTermDebtNoncurrent = getString(reportMap, "longTermDebtNoncurrent")
	report.ShortLongTermDebtTotal = getString(reportMap, "shortLongTermDebtTotal")
	report.OtherCurrentLiabilities = getString(reportMap, "otherCurrentLiabilities")
	report.OtherNonCurrentLiabilities = getString(reportMap, "otherNonCurrentLiabilities")
	report.TotalShareholderEquity = getString(reportMap, "totalShareholderEquity")
	report.TreasuryStock = getString(reportMap, "treasuryStock")
	report.RetainedEarnings = getString(reportMap, "retainedEarnings")
	report.CommonStock = getString(reportMap, "commonStock")
	report.CommonStockSharesOutstanding = getString(reportMap, "commonStockSharesOutstanding")

	return report
}
