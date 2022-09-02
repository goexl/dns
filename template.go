package uda

import (
	"context"
)

// 内部接口封装
// 使用模板方法设计模式
type template struct {
	aliyun executor
}

func (t *template) Add(ctx context.Context, domain string, rr string, value string, opts ...option) (err error) {
	options := defaultOptions
	for _, opt := range opts {
		opt.apply(options)
	}

	key := t.key(path, options.environment, options.separator)
	switch options.uoaType {
	case TypeCos:
		exist, err = t.cos.exist(ctx, key, options)
	}

	return
}
