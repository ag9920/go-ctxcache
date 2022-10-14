package goctxcache

import "context"

type keyType struct{}

var callCacheKey keyType

// WithCallCache 返回支持调用缓存的context
func WithCallCache(parent context.Context) context.Context {
	if parent.Value(callCacheKey) != nil {
		return parent
	}
	return context.WithValue(parent, callCacheKey, new(callCache))
}

// LoadFromCtxCache 从ctx中尝试获取key的缓存结果
// 如果不存在，调用loader;如果没有开启缓存，直接调用loader
func LoadFromCtxCache(ctx context.Context, key string, loader loadFunc) (interface{}, error) {
	var cacheItem *cacheItem
	v := ctx.Value(callCacheKey)
	if v == nil {
		cacheItem = nil
	} else {
		cacheItem = v.(*callCache).getOrCreateCacheItem(key)
	}

	// cache not enabled
	if cacheItem == nil {
		return loader(ctx)
	}

	// now that all routines hold references to the same cacheItem
	cacheItem.doOnce(ctx, loader)
	return cacheItem.ret, cacheItem.err
}
