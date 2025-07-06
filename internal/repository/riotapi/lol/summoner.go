package lol

import (
	"fmt"
	"net/http"

	"github.com/mgra04/backend-lol-v1/internal/repository/riotapi"
	"github.com/mgra04/backend-lol-v1/pkg/errors"

	jsoniter "github.com/json-iterator/go"
)

type SummonerDTO struct {
    Puuid     string `json:"puuid"`
    GameName  string `json:"gameName"`
    TagLine   string `json:"tagLine"`
}

func GetSummonerByRiotID(client *riotapi.RiotAPIClient, gameName, tagLine string) (*SummonerDTO, error) {
    url := fmt.Sprintf("https://europe.api.riotgames.com/riot/account/v1/accounts/by-riot-id/%s/%s?api_key=%s", gameName, tagLine, client.APIKey)
	println(url)

    client.RateLimiter.Wait(nil)

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return nil, errors.Wrap(err, "failed to create request", 500)
    }
    req.Header.Set("X-Riot-Token", client.APIKey)

    resp, err := client.HttpClient.Do(req)
    if err != nil {
        return nil, errors.Wrap(err, "riot api request failed", 502)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return nil, errors.New(fmt.Sprintf("riot api returned %d", resp.StatusCode), resp.StatusCode)
    }

    var result SummonerDTO
    if err := jsoniter.NewDecoder(resp.Body).Decode(&result); err != nil {
        return nil, errors.Wrap(err, "failed to decode riot api response", 500)
    }

    return &result, nil
}
