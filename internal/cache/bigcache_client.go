package cache

import (
	"time"

	"github.com/allegro/bigcache/v3"
)

type BigCache struct {
    Client *bigcache.BigCache
}

func NewBigCache() *BigCache {
    config := bigcache.Config{
        Shards:             1024,
        LifeWindow:         10 * time.Minute,
        CleanWindow:        5 * time.Minute,
        MaxEntriesInWindow: 1000 * 10 * 60,
        MaxEntrySize:       500,
        Verbose:            false,
        HardMaxCacheSize:   64, // MB
    }

    client, _ := bigcache.NewBigCache(config)
    return &BigCache{Client: client}
}
