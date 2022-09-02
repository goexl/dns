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
		secret gox.Secret
		typ    Type
		ttl    time.Duration
	}
)
