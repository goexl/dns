package uda

import (
	"time"

	"github.com/goexl/gox"
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		mode   mode
		secret gox.Secret
		typ    Type
		ttl    time.Duration
	}
)

func defaultOptions() *options {
	return &options{
		mode: modeAliyun,
		typ:  TypeA,
		ttl:  10 * time.Minute,
	}
}
