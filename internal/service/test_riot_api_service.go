package service

import (
	"fmt"

	"github.com/mgra04/backend-lol-v1/internal/constants"
	"github.com/mgra04/backend-lol-v1/internal/repository/riotapi"
	"github.com/rs/zerolog/log"
)

type TestRiotApiService struct {
	riotAPI riotapi.API
}

func NewTestRiotApiService(riotApi riotapi.API) *TestRiotApiService {
	return &TestRiotApiService {
		riotAPI: riotApi,
	}
}

type RiotApiProfileRawResponse struct {
	Account *riotapi.AccountDTO `json:"account"`
	Summoner *riotapi.SummonerDTO `json:"summoner"`
	TopChampionMasteries []riotapi.ChampionMasteryDTO `json:"topChampionMasteries"`
	TotalChampionMasteryScore int `json:"totalChampionMasteryScore"`
	Leagues []riotapi.LeagueEntryDTO `json:"leagues"`
	PlatformStatus *riotapi.PlatformDataDTO `json:"platformStatus"`
	MatchIds []string `json:"matchIds"`
	SingleMatch *riotapi.MatchDTO `json:"match"`
	SingleMatchTimeline *riotapi.MatchTimelineDTO `json:"matchTimeline"`
}

func (s *TestRiotApiService) LookupProfile(gameName, tagLine string, region constants.Region) (*RiotApiProfileRawResponse, error) {
	account, err := s.riotAPI.GetAccountByRiotID(gameName, tagLine, region)
	if err != nil {
		return nil, fmt.Errorf("failed to get account: %w", err)
	}

	summoner, err := s.riotAPI.GetSummonerByPUUID(account.Puuid, region)
	if err != nil {
		return nil, fmt.Errorf("failed to get summoner: %w", err)
	}

	topMasteries, err := s.riotAPI.GetTopChampionMasteriesByPUUID(account.Puuid, region)
	if err != nil {
		return nil, fmt.Errorf("failed to get top champion masteries: %w", err)
	}

	totalMastery, err := s.riotAPI.GetTotalChampionMasteryScoreByPUUID(account.Puuid, region)
	if err != nil {
		return nil, fmt.Errorf("failed to get total mastery score: %w", err)
	}

	leagues, err := s.riotAPI.GetLeaguesByPUUID(account.Puuid, region)
	if err != nil {
		return nil, fmt.Errorf("failed to get leagues: %w", err)
	}

	platformStatus, err := s.riotAPI.GetPlatformStatus(region)
	if err != nil {
		return nil, fmt.Errorf("failed to get platform status: %w", err)
	}

	const (
		start = 0
		count = 20
	)
	matchIds, err := s.riotAPI.GetMatchIDsByPUUID(account.Puuid, region, start, count)
	if err != nil {
		return nil, fmt.Errorf("failed to get match ids: %w", err)
	}
	log.Debug().Msgf("Fetched %d match IDs", len(matchIds))
	if len(matchIds) == 0 {
		return nil, fmt.Errorf("no matches found for player")
	}

	match, err := s.riotAPI.GetMatchByID(matchIds[0], region)
	if err != nil {
		return nil, fmt.Errorf("failed to get match info: %w", err)
	}
	log.Debug().Msg("Fetched match info")

	matchTimeline, err := s.riotAPI.GetMatchTimelineByMatchID(matchIds[0], region)
	if err != nil {
		return nil, fmt.Errorf("failed to get match timeline: %w", err)
	}
	log.Debug().Msg("Fetched match timeline")

	return &RiotApiProfileRawResponse{
		Account:              account,
		Summoner:             summoner,
		TopChampionMasteries: topMasteries,
		TotalChampionMasteryScore: totalMastery,
		Leagues:              leagues,
		PlatformStatus:       platformStatus,
		MatchIds: matchIds,
		SingleMatch: match,
		SingleMatchTimeline: matchTimeline,
	}, nil
}