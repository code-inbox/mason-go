package listener

import (
	"time"

	"github.com/aws/aws-sdk-go-v2/service/dynamodbstreams/types"
)

const (
	defaultBatchSize  = 100
	defaultInterval   = 5 * time.Second
	defaultRetryCount = 3
)

type Options struct {
	batchSize         int
	maxBatchWait      time.Duration
	debug             func(format string, args ...interface{})
	pollInterval      time.Duration
	shardIteratorType string
	retryCount        int
}

type Option func(*Options)

func WithBatchSize(n int) Option {
	return func(o *Options) {
		o.batchSize = n
	}
}

func WithDebug(fn func(format string, args ...interface{})) Option {
	return func(o *Options) {
		o.debug = fn
	}
}

func WithRetryCount(n int) Option {
	return func(o *Options) {
		o.retryCount = n
	}
}

func WithIteratorType(shardIteratorType string) Option {
	return func(o *Options) {
		o.shardIteratorType = shardIteratorType
	}
}

func WithPollInterval(interval time.Duration) Option {
	return func(o *Options) {
		o.pollInterval = interval
	}
}

func WithMaxBatchWait(interval time.Duration) Option {
	return func(o *Options) {
		o.maxBatchWait = interval
	}
}

func buildOptions(opts ...Option) Options {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.batchSize <= 0 || options.batchSize > 1000 {
		options.batchSize = defaultBatchSize
	}

	if options.debug == nil {
		options.debug = func(format string, args ...interface{}) {}
	}

	if options.pollInterval <= 0 {
		options.pollInterval = defaultInterval
	}

	if options.retryCount <= 0 {
		options.retryCount = defaultRetryCount
	}

	if options.maxBatchWait <= 0 {
		options.maxBatchWait = defaultInterval
	}

	if options.shardIteratorType == "" {
		options.shardIteratorType = string(types.ShardIteratorTypeLatest)
	}

	return options
}
