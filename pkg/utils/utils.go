package utils

import "github.com/mgra04/backend-lol-v1/internal/constants"

func GetRouteForRegion(region constants.Region) constants.Route {
    switch region {
    case constants.RegionEUW, constants.RegionEUNE, constants.RegionTR1, constants.RegionRU:
        return constants.RouteEurope
    case constants.RegionNA1, constants.RegionBR1, constants.RegionLA1, constants.RegionLA2:
        return constants.RouteAmericas
    case constants.RegionKR, constants.RegionJP1, constants.RegionOC1:
        return constants.RouteAsia
    default:
        return constants.RouteEurope
    }
}