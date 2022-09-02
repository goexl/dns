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
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		err = t.aliyun.add(ctx, domain, rr, value, _options)
	}

	return
}

func (t *template) Resolve(
	ctx context.Context,
	domain string, rr string, value string,
	opts ...option,
) (result *Result, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		result, err = t.aliyun.resolve(ctx, domain, rr, value, _options)
	}

	return
}

func (t *template) Get(ctx context.Context, domain string, rr string, opts ...option) (record *Record, err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		record, err = t.aliyun.get(ctx, domain, rr, _options)
	}

	return
}

func (t *template) Update(ctx context.Context, record *Record, value string, opts ...option) (err error) {
	_options := defaultOptions()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		err = t.aliyun.update(ctx, record, value, _options)
	}

	return
}
