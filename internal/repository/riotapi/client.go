package riotapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/mgra04/backend-lol-v1/internal/constants"
	"golang.org/x/time/rate"
)

const (
	apiURLFormat      = "%s://%s.%s%s"
	baseURL           = "api.riotgames.com"
	scheme            = "https"
	apiTokenHeaderKey = "X-Riot-Token"
)

type RiotApiClient struct {
	APIKey string
	HttpClient *http.Client
	RateLimiter *rate.Limiter
}

func NewRiotApiClient(key string) *RiotApiClient {
	return &RiotApiClient{
		APIKey: key,
        HttpClient: &http.Client{Timeout: 10 * time.Second},
        RateLimiter: rate.NewLimiter(rate.Every(time.Second), constants.RiotRateLimitBurst),
	}
}

func (c *RiotApiClient) getBaseURL(regionOrRoute string) string {
	return fmt.Sprintf("%s://%s.%s", scheme, regionOrRoute, baseURL)
}

func (c *RiotApiClient) Get(regionOrRoute string, path string, v interface{}) error {
	url := fmt.Sprintf("%s%s", c.getBaseURL(regionOrRoute), path)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}
	req.Header.Set(apiTokenHeaderKey, c.APIKey)

	resp, err := c.HttpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return mapAPIError(resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}