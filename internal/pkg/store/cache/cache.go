package cache

import "github.com/patrickmn/go-cache"

var (
	Cache *cache.Cache
)

func Init() error {
	Cache = cache.New(cache.NoExpiration, cache.NoExpiration)
	return nil
}
