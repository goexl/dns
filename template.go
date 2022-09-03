package uda

import (
	"context"
)

// 内部接口封装
// 使用模板方法设计模式
type template struct {
	aliyun executor

	options *options
}

func newTemplate(options *options) *template {
	return &template{
		aliyun: newAliyun(),

		options: options,
	}
}

func (t *template) Add(ctx context.Context, domain string, subdomain string, value string, opts ...option) (err error) {
	_options := t.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		err = t.aliyun.add(ctx, domain, subdomain, value, _options)
	}

	return
}

func (t *template) Resolve(
	ctx context.Context,
	domain string, subdomain string, value string,
	opts ...option,
) (original string, do bool, err error) {
	_options := t.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if record, getErr := t.Get(ctx, domain, subdomain, opts...); nil != getErr {
		err = getErr
	} else if nil != record {
		original = record.Value
		if value != record.Value {
			do = true
			err = t.Update(ctx, record, value, opts...)
		}
	} else {
		do = true
		err = t.Add(ctx, domain, subdomain, value, opts...)
	}

	return
}

func (t *template) Get(ctx context.Context, domain string, subdomain string, opts ...option) (record *Record, err error) {
	_options := t.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		record, err = t.aliyun.get(ctx, domain, subdomain, _options)
	}

	return
}

func (t *template) Update(ctx context.Context, record *Record, value string, opts ...option) (err error) {
	_options := t.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		err = t.aliyun.update(ctx, record, value, _options)
	}

	return
}
