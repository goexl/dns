package dns

import (
	"time"
)

var (
	_        = Aliyun
	_ option = (*optionTtl)(nil)
)

type optionTtl struct {
	ttl time.Duration
}

// Ttl 生存时间
func Ttl(ttl time.Duration) *optionTtl {
	return &optionTtl{
		ttl: ttl,
	}
}

// TTL 生存时间
func TTL(ttl time.Duration) *optionTtl {
	return Ttl(ttl)
}

func (t *optionTtl) apply(options *options) {
	options.ttl = t.ttl
}
