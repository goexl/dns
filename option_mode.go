package uda

var (
	_        = Aliyun
	_ option = (*optionMode)(nil)
)

type optionMode struct {
	mode mode
}

// Aliyun 阿里云
func Aliyun() *optionMode {
	return &optionMode{
		mode: modeAliyun,
	}
}

func (m *optionMode) apply(options *options) {
	options.mode = m.mode
}
