package riotapi

import "github.com/mgra04/backend-lol-v1/internal/constants"

type API interface {
	GetAccountByRiotID(gameName string, tagLine string, region constants.Region) (*AccountDTO, error)
	GetSummonerByPUUID(puuid string, region constants.Region) (*SummonerDTO, error)
	GetTopChampionMasteriesByPUUID(puuid string, region constants.Region) ([]ChampionMasteryDTO, error)
	GetTotalChampionMasteryScoreByPUUID(puuid string, region constants.Region) (int, error)
	GetLeaguesByPUUID(puuid string, region constants.Region) ([]LeagueEntryDTO, error)
	GetPlatformStatus(region constants.Region) (*PlatformDataDTO, error)
	GetMatchIDsByPUUID(puuid string, region constants.Region, start int, count int) ([]string, error)
	GetMatchByID(matchID string, region constants.Region) (*MatchDTO, error)
	GetMatchTimelineByMatchID(matchID string, region constants.Region) (*TimelineDTO, error)
}
