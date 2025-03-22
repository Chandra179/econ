package fundamental

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

	// Build request URL
	url := fmt.Sprintf("%s/query?function=INCOME_STATEMENT&symbol=%s&apikey=%s",
		cfg.AlphaVantageBaseURL,
		params.Symbol,
		cfg.AlphaVantageAPIKey)

	// Make the request
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error making request to Alpha Vantage: %w", err)
	}
	defer resp.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	// Check for error messages in the response
	var errorResponse map[string]interface{}
	if err := json.Unmarshal(body, &errorResponse); err == nil {
		if errorMsg, exists := errorResponse["Error Message"]; exists {
			return nil, fmt.Errorf("API error: %v", errorMsg)
		}
		if errorMsg, exists := errorResponse["Information"]; exists {
			return nil, fmt.Errorf("API information: %v", errorMsg)
		}
	}

	// Parse the response
	var result IncomeStatementResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("error parsing income statement data: %w", err)
	}

	return &result, nil
}
