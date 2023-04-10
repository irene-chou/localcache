package localcache

import (
	"sync"
	"time"
)

const (
	EXPIRATION_TTL = 30
)

var (
	timeNow = time.Now
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type cacheItem struct {
	value      interface{}
	expiryTime time.Time
}

type localCache struct {
	items map[string]*cacheItem
	mutex sync.RWMutex
}

func (lc *localCache) Get(key string) (interface{}, bool) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	item, ok := lc.items[key]
	if !ok {
		return nil, false
	}
	if item.expiryTime.Before(timeNow()) {
		delete(lc.items, key)
		return nil, false
	}
	return item.value, true
}

func (lc *localCache) Set(key string, value interface{}) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	expiryTime := timeNow().Add(EXPIRATION_TTL * time.Second)
	lc.items[key] = &cacheItem{value: value, expiryTime: expiryTime}
}

func New() Cache {
	return &localCache{items: make(map[string]*cacheItem)}
}
