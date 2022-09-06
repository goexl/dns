package dns

var (
	_        = Aliyun
	_        = TencentCloud
	_ option = (*optionProvider)(nil)
)

type optionProvider struct {
	provider provider
}

// Provider 厂商
func Provider(provider provider) *optionProvider {
	return &optionProvider{
		provider: provider,
	}
}

// Aliyun 阿里云
func Aliyun() *optionProvider {
	return Provider(providerAliyun)
}

// TencentCloud 腾讯云
func TencentCloud() *optionProvider {
	return Provider(providerTencentCloud)
}

func (p *optionProvider) apply(options *options) {
	options.provider = p.provider
}
