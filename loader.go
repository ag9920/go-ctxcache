package goctxcache

import "context"

type loadFunc func(context.Context) (interface{}, error)
