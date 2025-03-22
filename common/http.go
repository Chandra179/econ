package common

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// BuildRequestURL builds the complete URL with query parameters
func BuildRequestURL(baseURL string, params map[string]string) (string, error) {
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

// MakeAPIRequest performs the HTTP request to an API and returns the parsed JSON response
func MakeAPIRequest(baseURL string, params map[string]string) (map[string]interface{}, error) {
	client := &http.Client{}

	// Build URL with query parameters
	fullURL, err := BuildRequestURL(baseURL, params)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest("GET", fullURL, nil)
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

	// Check for API error
	if errorMsg, ok := result["Error Message"].(string); ok {
		return nil, fmt.Errorf("API error: %s", errorMsg)
	}

	return result, nil
}
