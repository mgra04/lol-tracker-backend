package constants

import "time"

// HTTP Server
const (
    ReadTimeout  = 15 * time.Second
    WriteTimeout = 15 * time.Second
    IdleTimeout  = 60 * time.Second
)

// Riot API limits
const (
    RiotRateLimitPerSecond = 10
    RiotRateLimitBurst     = 20
)

// All existing regions (riot api)
type Region string

const (
	RegionBR1           Region = "br1"
	RegionEUNE   Region = "eun1"
	RegionEUW       Region = "euw1"
	RegionJP1            Region = "jp1"
	RegionKR            Region = "kr"
	RegionLA1 Region = "la1"
	RegionLA2 Region = "la2"
	RegionME1       Region = "me1"
	RegionNA1      Region = "na1"
	RegionOC1           Region = "oc1"
	RegionPBE               Region = "pbe1"
	RegionSG2 Region = "sg2"
	RegionRU      Region = "ru"
	RegionTR1   Region = "tr1"
	RegionTW2   Region = "tw2"
	RegionVN2  Region = "vn2"
)


// All existing routes (riotapi)
type Route string

const (
	RouteAmericas Route = "americas"
	RouteAsia     Route = "asia"
	RouteEurope   Route = "europe"
	RouteSEA      Route = "sea"
)

// All used lol endpoints