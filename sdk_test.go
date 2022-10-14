package goctxcache

import (
	"context"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func randFunc(ctx context.Context, flag string) (int, error) {
	return rand.Int(), nil
}

func RandFunc(ctx context.Context, flag string) (int, error) {
	key := "RandFunc:" + flag
	r, e := LoadFromCtxCache(ctx, key, func(ctx context.Context) (interface{}, error) {
		return randFunc(ctx, flag)
	})
	return r.(int), e
}

func TestLoadFromCtxCacheSuccess(t *testing.T) {
	assert := assert.New(t)
	ctx := context.Background()

	var ret1, ret2 int
	var err1, err2 error
	const flag = ""

	// without cache, random return
	ret1, err1 = RandFunc(ctx, flag)
	assert.NoError(err1)
	ret2, err2 = RandFunc(ctx, flag)
	assert.NoError(err2)
	assert.NotEqual(ret1, ret2)

	// with cache, same return
	cacheCtx := WithCallCache(ctx)
	ret1, err1 = RandFunc(cacheCtx, flag)
	assert.NoError(err1)
	ret2, err2 = RandFunc(cacheCtx, flag)
	assert.NoError(err2)
	assert.Equal(ret1, ret2)
}
