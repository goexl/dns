package uda

import (
	"github.com/goexl/gox"
)

type (
	option interface {
		apply(options *options)
	}

	options struct {
		secret gox.Secret
	}
)
