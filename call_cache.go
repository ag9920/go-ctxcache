package goctxcache

import "sync"

type callCache struct {
	lock sync.RWMutex
	m    map[string]*cacheItem
}

// getOrCreateCacheItem 从callCache中获取指定key的cacheItem(不存在则创建一个)。保证并发安全, 不会返回nil
func (cache *callCache) getOrCreateCacheItem(key string) *cacheItem {
	cache.lock.RLock()
	cr, ok := cache.m[key]
	cache.lock.RUnlock()
	if ok {
		return cr
	}

	cache.lock.Lock()
	defer cache.lock.Unlock()
	if cache.m == nil {
		cache.m = make(map[string]*cacheItem)
	} else {
		cr, ok = cache.m[key]
	}
	if !ok {
		cr = &cacheItem{}
		cache.m[key] = cr
	}
	return cr
}
