package gbapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

var BaseURL = "http://localhost:9090"

type RestClient struct {
	BaseURL    string
	HTTPClient *http.Client
	Headers    map[string]string
}

func NewRestClient(baseURL string, headers map[string]string) *RestClient {
	return &RestClient{
		BaseURL: baseURL,
		Headers: headers,
		HTTPClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// doRequest handles the actual HTTP request
func (c *RestClient) doRequest(method, path string, body interface{}, response interface{}) error {
	url := c.BaseURL + path

	var bodyReader io.Reader

	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal body: %w", err)
		}
		bodyReader = bytes.NewBuffer(jsonData)
	}
	req, err := http.NewRequest(method, url, bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	for k, v := range c.Headers {
		req.Header.Set(k, v)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		bodyBytes, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%d response: %s", resp.StatusCode, string(bodyBytes))
	}

	if response != nil {
		return json.NewDecoder(resp.Body).Decode(response)
	}
	return nil
}

// Get performs a GET request
func (c *RestClient) Get(path string, response interface{}) error {
	return c.doRequest(http.MethodGet, path, nil, response)
}

// Post performs a POST request
func (c *RestClient) Post(path string, body interface{}, response interface{}) error {
	return c.doRequest(http.MethodPost, path, body, response)
}

// Put performs a PUT request
func (c *RestClient) Put(path string, body interface{}, response interface{}) error {
	return c.doRequest(http.MethodPut, path, body, response)
}

// Delete performs a DELETE request
func (c *RestClient) Delete(path string, response interface{}) error {
	return c.doRequest(http.MethodDelete, path, nil, response)
}

// Patch performs a DELETE request
func (c *RestClient) Patch(path string, body interface{}, response interface{}) error {
	return c.doRequest(http.MethodPatch, path, body, response)
}
