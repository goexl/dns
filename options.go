package dns

var _ = NewOptions

type Options []option

// NewOptions 创建选项
func NewOptions(opts ...option) Options {
	return opts
}

func (o *Options) Add(opts ...option) {
	*o = append(*o, opts...)
}
