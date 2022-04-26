package dingding

import (
	"time"
)

const (
	DefaultRequestTimeout  = time.Second * 30 // 30s
	DefaultWriteBufferSize = 1024 * 1024 * 20 // 20 MB
	DefaultReadBufferSize  = 1024 * 1024 * 20 // 20 MB
)

type Options struct {
	// 钉钉接口应用的 appKey
	appKey string
	// 钉钉接口应用的 appSecret
	appSecret string

	// access_token
	token string

	// 请求超时时间
	timeout time.Duration

	writeSize int
	readSize  int
}

func NewOptions(opts ...Option) Options {
	options := Options{}

	for _, opt := range opts {
		opt(&options)
	}

	if options.timeout == 0 {
		options.timeout = DefaultRequestTimeout
	}
	if options.writeSize == 0 {
		options.writeSize = DefaultWriteBufferSize
	}
	if options.readSize == 0 {
		options.readSize = DefaultReadBufferSize
	}

	return options
}

type Option func(*Options)

func WithAppKey(appKey string) Option {
	return func(o *Options) {
		o.appKey = appKey
	}
}

func WithAppSecret(appSecret string) Option {
	return func(o *Options) {
		o.appSecret = appSecret
	}
}

func WithToken(token string) Option {
	return func(o *Options) {
		o.token = token
	}
}

func WithTimeout(timeout time.Duration) Option {
	return func(o *Options) {
		o.timeout = timeout
	}
}

func WithWriteSize(size int) Option {
	return func(o *Options) {
		o.writeSize = size
	}
}

func WithReadSize(size int) Option {
	return func(o *Options) {
		o.readSize = size
	}
}
