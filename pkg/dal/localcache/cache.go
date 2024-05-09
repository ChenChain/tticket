package localcache

import (
	gocache "github.com/patrickmn/go-cache"
	"time"
)

var cache *gocache.Cache

func init() {
	cache = gocache.New(10*time.Minute, 20*time.Minute)
}

func Cache() *gocache.Cache {
	return cache
}
