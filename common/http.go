package common

import (
	"io"
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

// GetAPIRequest performs the Get HTTP request to an API and returns the parsed JSON response
func GetAPIRequest(baseURL string, params map[string]string) (io.ReadCloser, error) {
	return MakeAPIRequest(baseURL, params, "GET")
}

// MakeAPIRequest performs the HTTP request to an API and returns the parsed JSON response
func MakeAPIRequest(baseURL string, params map[string]string, apiMethod string) (io.ReadCloser, error) {
	client := &http.Client{}

	// Build URL with query parameters
	fullURL, err := BuildRequestURL(baseURL, params)
	if err != nil {
		return nil, err
	}

	// Create request
	req, err := http.NewRequest(apiMethod, fullURL, nil)
	if err != nil {
		return nil, err
	}

	// Execute request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp.Body, nil
}
