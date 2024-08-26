package utils

import (
	"bitroom/constants"
	"sync"

	"github.com/patrickmn/go-cache"
)

var (
	cacheInstance *cache.Cache
	cacheOnce     sync.Once
)

func GetCache() *cache.Cache {
	cacheOnce.Do(func() {
		cacheInstance = cache.New(constants.CacheItemTimeExpiration, constants.CachePurgesTimeExpiratoin)
	})
	return cacheInstance
}
