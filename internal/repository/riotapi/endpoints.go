package riotapi

import (
	"fmt"

	"github.com/mgra04/backend-lol-v1/internal/constants"
	"github.com/mgra04/backend-lol-v1/pkg/utils"
)

func (c *RiotApiClient) GetAccountByRiotID(gameName string, tagLine string, region constants.Region) (*AccountDTO, error) {
	route := utils.GetRouteForRegion(region)
	var account AccountDTO
	path := fmt.Sprintf("/riot/account/v1/accounts/by-riot-id/%s/%s", gameName, tagLine)
	err := c.Get(string(route), path, &account)
	return &account, err
}

func (c *RiotApiClient) GetSummonerByPUUID(puuid string, region constants.Region) (*SummonerDTO, error) {
	var summoner SummonerDTO
	path := fmt.Sprintf("/lol/summoner/v4/summoners/by-puuid/%s", puuid)
	err := c.Get(string(region), path, &summoner)
	return &summoner, err
}

func (c *RiotApiClient) GetTopChampionMasteriesByPUUID(puuid string, region constants.Region) ([]ChampionMasteryDTO, error) {
	var masteries []ChampionMasteryDTO
	path := fmt.Sprintf("/lol/champion-mastery/v4/champion-masteries/by-puuid/%s/top?count=12", puuid)
	err := c.Get(string(region), path, &masteries)
	return masteries, err
}

func (c *RiotApiClient) GetTotalChampionMasteryScoreByPUUID(puuid string, region constants.Region) (int, error) {
	var TotalChampionMasteryScore int
	path := fmt.Sprintf("/lol/champion-mastery/v4/scores/by-puuid/%s", puuid)
	err := c.Get(string(region), path, &TotalChampionMasteryScore)
	return TotalChampionMasteryScore, err
}

func (c *RiotApiClient) GetLeaguesByPUUID(puuid string, region constants.Region) ([]LeagueEntryDTO, error) {
	var leagues []LeagueEntryDTO
	path := fmt.Sprintf("/lol/league/v4/entries/by-puuid/%s", puuid)
	err := c.Get(string(region), path, &leagues)
	return leagues, err
}

func (c *RiotApiClient) GetPlatformStatus(region constants.Region) (*PlatformDataDTO, error) {
	var status PlatformDataDTO
	path := "/lol/status/v4/platform-data"
	err := c.Get(string(region), path, &status)
	return &status, err
}

func (c *RiotApiClient) GetMatchIDsByPUUID(puuid string, region constants.Region, start int, count int) ([]string, error) {
	route := utils.GetRouteForRegion(region)
	var ids []string
	path := fmt.Sprintf("/lol/match/v5/matches/by-puuid/%s/ids?start=%d&count=%d", puuid, start, count)
	err := c.Get(string(route), path, &ids)
	return ids, err
}

func (c *RiotApiClient) GetMatchByID(matchID string, region constants.Region) (*MatchDTO, error) {
	route := utils.GetRouteForRegion(region)
	var match MatchDTO
	path := fmt.Sprintf("/lol/match/v5/matches/%s", matchID)
	err := c.Get(string(route), path, &match)
	return &match, err
}

func (c *RiotApiClient) GetMatchTimelineByMatchID(matchID string, region constants.Region) (*MatchTimelineDTO, error) {
	route := utils.GetRouteForRegion(region)
	var timeline MatchTimelineDTO
	path := fmt.Sprintf("/lol/match/v5/matches/%s/timeline", matchID)
	err := c.Get(string(route), path, &timeline)
	return &timeline, err
}