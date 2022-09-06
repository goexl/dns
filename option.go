package dns

import (
	"time"

	"github.com/goexl/gox"
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		provider provider
		secret   gox.Secret
		typ      Type
		ttl      time.Duration
	}
)

func defaultOptions() *options {
	return &options{
		provider: providerAliyun,
		typ:      TypeA,
		ttl:      10 * time.Minute,
	}
}

func (o *options) clone() *options {
	return &options{
		provider: o.provider,
		typ:      o.typ,
		ttl:      o.ttl,
	}
}
