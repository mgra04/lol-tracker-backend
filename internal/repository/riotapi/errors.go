package riotapi

import "fmt"

func mapAPIError(statusCode int) error {
	switch statusCode {
	case 400:
		return fmt.Errorf("bad request to Riot API")
	case 401:
		return fmt.Errorf("unauthorized - check API key")
	case 403:
		return fmt.Errorf("forbidden - check API key or permissions")
	case 404:
		return fmt.Errorf("resource not found")
	case 429:
		return fmt.Errorf("rate limit exceeded")
	case 500, 502, 503, 504:
		return fmt.Errorf("riot API server error (%d)", statusCode)
	default:
		return fmt.Errorf("unexpected riot API error (%d)", statusCode)
	}
}
