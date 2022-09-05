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

func (o *options) clone() *options {
	return &options{
		mode: o.mode,
		typ:  o.typ,
		ttl:  o.ttl,
	}
}
