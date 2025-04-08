package fundamental

import (
	"encoding/json"
	"stock/common"
	"stock/config"
)

// IncomeStatementParams holds parameters for retrieving income statement data
type IncomeStatementParams struct {
	Symbol string
}

// IncomeStatementResponse defines the structure for income statement data
type IncomeStatementResponse struct {
	Symbol           string                  `json:"symbol"`
	AnnualReports    []IncomeStatementReport `json:"annualReports"`
	QuarterlyReports []IncomeStatementReport `json:"quarterlyReports"`
}

// IncomeStatementReport represents a single income statement report
type IncomeStatementReport struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	GrossProfit                       string `json:"grossProfit"`
	TotalRevenue                      string `json:"totalRevenue"`
	CostOfRevenue                     string `json:"costOfRevenue"`
	CostofGoodsAndServicesSold        string `json:"costofGoodsAndServicesSold"`
	OperatingIncome                   string `json:"operatingIncome"`
	SellingGeneralAndAdministrative   string `json:"sellingGeneralAndAdministrative"`
	ResearchAndDevelopment            string `json:"researchAndDevelopment"`
	OperatingExpenses                 string `json:"operatingExpenses"`
	InvestmentIncomeNet               string `json:"investmentIncomeNet"`
	NetInterestIncome                 string `json:"netInterestIncome"`
	InterestIncome                    string `json:"interestIncome"`
	InterestExpense                   string `json:"interestExpense"`
	NonInterestIncome                 string `json:"nonInterestIncome"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	Depreciation                      string `json:"depreciation"`
	DepreciationAndAmortization       string `json:"depreciationAndAmortization"`
	IncomeBeforeTax                   string `json:"incomeBeforeTax"`
	IncomeTaxExpense                  string `json:"incomeTaxExpense"`
	InterestAndDebtExpense            string `json:"interestAndDebtExpense"`
	NetIncomeFromContinuingOperations string `json:"netIncomeFromContinuingOperations"`
	ComprehensiveIncomeNetOfTax       string `json:"comprehensiveIncomeNetOfTax"`
	Ebit                              string `json:"ebit"`
	Ebitda                            string `json:"ebitda"`
	NetIncome                         string `json:"netIncome"`
}

// GetIncomeStatement retrieves income statement data for a given symbol
func GetIncomeStatement(params IncomeStatementParams) (*IncomeStatementResponse, error) {
	cfg := config.GetConfig()

	// Building query parameters
	queryParams := map[string]string{
		"function": "INCOME_STATEMENT",
		"symbol":   params.Symbol,
		"apikey":   cfg.AlphaVantageAPIKey,
	}
	// Make HTTP request
	respBody, err := common.GetAPIRequest(cfg.AlphaVantageBaseURL, queryParams)
	if err != nil {
		return nil, err
	}
	defer respBody.Close()

	// Parse the response
	result := &IncomeStatementResponse{
		Symbol: params.Symbol,
	}
	if err := json.NewDecoder(respBody).Decode(result); err != nil {
		return nil, err
	}

	return result, nil
}
