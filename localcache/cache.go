package localcache

import (
	"sync"
	"time"
)

const (
	EXPIRATION_TTL = 30
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

type localCache struct {
	cacheMap map[string]interface{}
	mutex    sync.RWMutex
}

func (lc *localCache) Get(key string) (interface{}, bool) {
	lc.mutex.RLock()
	defer lc.mutex.RUnlock()

	value, ok := lc.cacheMap[key]
	return value, ok
}

func (lc *localCache) Set(key string, value interface{}) {
	lc.mutex.Lock()
	defer lc.mutex.Unlock()

	lc.cacheMap[key] = value
	time.AfterFunc(EXPIRATION_TTL*time.Second, func() {
		lc.mutex.Lock()
		defer lc.mutex.Unlock()
		delete(lc.cacheMap, key)
	})
}

func New() Cache {
	return &localCache{
		cacheMap: make(map[string]interface{}),
	}
}
