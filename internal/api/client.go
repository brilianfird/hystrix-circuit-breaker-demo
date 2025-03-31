package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client handles API operations
type Client struct {
	httpClient *http.Client
}

// NewClient creates a new instance of API Client
func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// CallMockEndpoint makes a request to the mock API endpoint
func (c *Client) CallMockEndpoint() (string, error) {
	resp, err := c.httpClient.Get("http://localhost:8080/mock-api")
	if err != nil {
		return "", fmt.Errorf("failed to call mock API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("received error status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read response body: %w", err)
	}

	return string(body), nil
}
