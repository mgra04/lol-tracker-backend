package riotapi

import (
	"net/http"
	"time"

	"github.com/mgra04/backend-lol-v1/internal/constants"
	"golang.org/x/time/rate"
)

type RiotAPIClient struct {
    APIKey     string
    HttpClient *http.Client
    RateLimiter *rate.Limiter
}

func NewRiotApiClient(apiKey string) *RiotAPIClient {
    return &RiotAPIClient{
        APIKey: apiKey,
        HttpClient: &http.Client{Timeout: 10 * time.Second},
        RateLimiter: rate.NewLimiter(rate.Every(time.Second), constants.RiotRateLimitBurst),
    }
}