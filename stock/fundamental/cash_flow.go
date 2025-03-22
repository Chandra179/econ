package fundamental

import (
	"stock/common"
	"stock/config"
)

// CashFlowParams holds parameters for retrieving cash flow data
type CashFlowParams struct {
	Symbol string // Required: Stock symbol (e.g., IBM)
}

// CashFlowResponse defines the response format for cash flow data
type CashFlowResponse struct {
	Symbol           string           `json:"symbol"`
	AnnualReports    []CashFlowReport `json:"annualReports"`
	QuarterlyReports []CashFlowReport `json:"quarterlyReports"`
}

// CashFlowReport represents a single cash flow report
type CashFlowReport struct {
	FiscalDateEnding                                          string `json:"fiscalDateEnding"`
	ReportedCurrency                                          string `json:"reportedCurrency"`
	OperatingCashflow                                         string `json:"operatingCashflow"`
	PaymentsForOperatingActivities                            string `json:"paymentsForOperatingActivities"`
	ProceedsFromOperatingActivities                           string `json:"proceedsFromOperatingActivities"`
	ChangeInOperatingLiabilities                              string `json:"changeInOperatingLiabilities"`
	ChangeInOperatingAssets                                   string `json:"changeInOperatingAssets"`
	DepreciationDepletionAndAmortization                      string `json:"depreciationDepletionAndAmortization"`
	CapitalExpenditures                                       string `json:"capitalExpenditures"`
	ChangeInReceivables                                       string `json:"changeInReceivables"`
	ChangeInInventory                                         string `json:"changeInInventory"`
	ProfitLoss                                                string `json:"profitLoss"`
	CashflowFromInvestment                                    string `json:"cashflowFromInvestment"`
	CashflowFromFinancing                                     string `json:"cashflowFromFinancing"`
	ProceedsFromRepaymentsOfShortTermDebt                     string `json:"proceedsFromRepaymentsOfShortTermDebt"`
	PaymentsForRepurchaseOfCommonStock                        string `json:"paymentsForRepurchaseOfCommonStock"`
	PaymentsForRepurchaseOfEquity                             string `json:"paymentsForRepurchaseOfEquity"`
	PaymentsForRepurchaseOfPreferredStock                     string `json:"paymentsForRepurchaseOfPreferredStock"`
	DividendPayout                                            string `json:"dividendPayout"`
	DividendPayoutCommonStock                                 string `json:"dividendPayoutCommonStock"`
	DividendPayoutPreferredStock                              string `json:"dividendPayoutPreferredStock"`
	ProceedsFromIssuanceOfCommonStock                         string `json:"proceedsFromIssuanceOfCommonStock"`
	ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet string `json:"proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet"`
	ProceedsFromIssuanceOfPreferredStock                      string `json:"proceedsFromIssuanceOfPreferredStock"`
	ProceedsFromRepurchaseOfEquity                            string `json:"proceedsFromRepurchaseOfEquity"`
	ProceedsFromSaleOfTreasuryStock                           string `json:"proceedsFromSaleOfTreasuryStock"`
	ChangeInCashAndCashEquivalents                            string `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate                                      string `json:"changeInExchangeRate"`
	NetIncome                                                 string `json:"netIncome"`
}

// GetCashFlow fetches cash flow data from Alpha Vantage API
func GetCashFlow(params CashFlowParams) (*CashFlowResponse, error) {
	// Get API configuration
	cfg := config.GetConfig()

	// Building query parameters
	queryParams := map[string]string{
		"function": "CASH_FLOW",
		"symbol":   params.Symbol,
		"apikey":   cfg.AlphaVantageAPIKey,
	}

	// Make HTTP request
	resp, err := common.MakeAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}

	// Convert generic response to CashFlowResponse
	cashFlowResp := &CashFlowResponse{
		Symbol: params.Symbol,
	}

	// Extract annual reports
	if annualReports, ok := resp["annualReports"].([]interface{}); ok {
		for _, report := range annualReports {
			if reportMap, ok := report.(map[string]interface{}); ok {
				cashFlowReport := mapToCashFlowReport(reportMap)
				cashFlowResp.AnnualReports = append(cashFlowResp.AnnualReports, cashFlowReport)
			}
		}
	}

	// Extract quarterly reports
	if quarterlyReports, ok := resp["quarterlyReports"].([]interface{}); ok {
		for _, report := range quarterlyReports {
			if reportMap, ok := report.(map[string]interface{}); ok {
				cashFlowReport := mapToCashFlowReport(reportMap)
				cashFlowResp.QuarterlyReports = append(cashFlowResp.QuarterlyReports, cashFlowReport)
			}
		}
	}

	return cashFlowResp, nil
}

// mapToCashFlowReport converts a map to a CashFlowReport struct
func mapToCashFlowReport(reportMap map[string]interface{}) CashFlowReport {
	report := CashFlowReport{}

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
	report.OperatingCashflow = getString(reportMap, "operatingCashflow")
	report.PaymentsForOperatingActivities = getString(reportMap, "paymentsForOperatingActivities")
	report.ProceedsFromOperatingActivities = getString(reportMap, "proceedsFromOperatingActivities")
	report.ChangeInOperatingLiabilities = getString(reportMap, "changeInOperatingLiabilities")
	report.ChangeInOperatingAssets = getString(reportMap, "changeInOperatingAssets")
	report.DepreciationDepletionAndAmortization = getString(reportMap, "depreciationDepletionAndAmortization")
	report.CapitalExpenditures = getString(reportMap, "capitalExpenditures")
	report.ChangeInReceivables = getString(reportMap, "changeInReceivables")
	report.ChangeInInventory = getString(reportMap, "changeInInventory")
	report.ProfitLoss = getString(reportMap, "profitLoss")
	report.CashflowFromInvestment = getString(reportMap, "cashflowFromInvestment")
	report.CashflowFromFinancing = getString(reportMap, "cashflowFromFinancing")
	report.ProceedsFromRepaymentsOfShortTermDebt = getString(reportMap, "proceedsFromRepaymentsOfShortTermDebt")
	report.PaymentsForRepurchaseOfCommonStock = getString(reportMap, "paymentsForRepurchaseOfCommonStock")
	report.PaymentsForRepurchaseOfEquity = getString(reportMap, "paymentsForRepurchaseOfEquity")
	report.PaymentsForRepurchaseOfPreferredStock = getString(reportMap, "paymentsForRepurchaseOfPreferredStock")
	report.DividendPayout = getString(reportMap, "dividendPayout")
	report.DividendPayoutCommonStock = getString(reportMap, "dividendPayoutCommonStock")
	report.DividendPayoutPreferredStock = getString(reportMap, "dividendPayoutPreferredStock")
	report.ProceedsFromIssuanceOfCommonStock = getString(reportMap, "proceedsFromIssuanceOfCommonStock")
	report.ProceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet = getString(reportMap, "proceedsFromIssuanceOfLongTermDebtAndCapitalSecuritiesNet")
	report.ProceedsFromIssuanceOfPreferredStock = getString(reportMap, "proceedsFromIssuanceOfPreferredStock")
	report.ProceedsFromRepurchaseOfEquity = getString(reportMap, "proceedsFromRepurchaseOfEquity")
	report.ProceedsFromSaleOfTreasuryStock = getString(reportMap, "proceedsFromSaleOfTreasuryStock")
	report.ChangeInCashAndCashEquivalents = getString(reportMap, "changeInCashAndCashEquivalents")
	report.ChangeInExchangeRate = getString(reportMap, "changeInExchangeRate")
	report.NetIncome = getString(reportMap, "netIncome")

	return report
}
