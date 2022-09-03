package uda

// New 创建域名解析适配器
func New(opts ...option) (resolver Resolver) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}
	resolver = newTemplate(_options)

	return
}
