package uda

import (
	"github.com/goexl/gox"
)

var (
	_        = Secret
	_ option = (*optionSecret)(nil)
)

type optionSecret struct {
	secret gox.Secret
}

// Secret 授权
func Secret(ak string, sk string) *optionSecret {
	return &optionSecret{
		secret: gox.Secret{
			Ak: ak,
			Sk: sk,
		},
	}
}

func (s *optionSecret) apply(options *options) {
	options.secret = s.secret
}
