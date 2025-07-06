package service

import (
	"github.com/mgra04/backend-lol-v1/internal/cache"
	"github.com/mgra04/backend-lol-v1/internal/repository/mongo"
	"github.com/mgra04/backend-lol-v1/internal/repository/riotapi"
	"github.com/mgra04/backend-lol-v1/internal/repository/riotapi/lol"
	"github.com/mgra04/backend-lol-v1/pkg/errors"
)

type ProfileService struct {
    riotClient *riotapi.RiotAPIClient
    mongoRepo  *mongo.Repository
    cache      *cache.BigCache
}

func NewProfileService(
    riot *riotapi.RiotAPIClient,
    mongoRepo *mongo.Repository,
    cache *cache.BigCache,
) *ProfileService {
    return &ProfileService{
        riotClient: riot,
        mongoRepo:  mongoRepo,
        cache:      cache,
    }
}

type ProfileDTO struct {
    Puuid        string `json:"puuid"`
    RiotUserName string `json:"riotUserName"`
    RiotTagLine  string `json:"riotTagLine"`
}

// GetProfile fetches the PUUID of a summoner by Riot ID
func (s *ProfileService) GetProfile(gameName string, tagLine string) (*ProfileDTO, error) {
    // Call Riot API for account info
    summoner, err := lol.GetSummonerByRiotID(s.riotClient, gameName, tagLine)
    if err != nil {
        return nil, errors.Wrap(err, "failed to fetch summoner from Riot API", 500)
    }

    profile := &ProfileDTO{
        Puuid:        summoner.Puuid,
        RiotUserName: summoner.GameName,
        RiotTagLine:  summoner.GameName,
    }

    // (TODO): Save to MongoDB and cache in BigCache
    return profile, nil
}