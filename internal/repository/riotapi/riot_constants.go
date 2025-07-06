package riotapi

type Region string
type Route string

const (
	RegionBR1  Region = "br1"
	RegionEUNE Region = "eun1"
	RegionEUW  Region = "euw1"
	RegionNA   Region = "na1"
	RegionKR   Region = "kr"
	RegionJP   Region = "jp1"
	RegionOCE  Region = "oc1"

	RouteAmericas Route = "americas"
	RouteEurope   Route = "europe"
	RouteAsia     Route = "asia"
)
