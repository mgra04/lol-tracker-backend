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
