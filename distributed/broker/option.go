package broker

import (
	"context"

	"github.com/lingyao2333/go-basics-utils/options"
)

func SetRetry(retry bool) options.Option {
	return func(c interface{}) {
		c.(*Broker).retry = retry
	}
}

func SetRetryFn(fn func(ctx context.Context)) options.Option {
	return func(c interface{}) {
		c.(*Broker).retryFn = fn
	}
}
