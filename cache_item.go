package goctxcache

import (
	"context"
	"sync"
)

type cacheItem struct {
	ret  interface{}
	err  error
	once sync.Once
}

func (ci *cacheItem) doOnce(ctx context.Context, loader loadFunc) {
	ci.once.Do(func() {
		ci.ret, ci.err = loader(ctx)
	})
}
