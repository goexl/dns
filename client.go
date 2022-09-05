package dns

import (
	"context"
)

var _ = New

// Client 客户端
type Client struct {
	aliyun executor

	options *options
}

// New 创建客户端
func New(opts ...option) (client *Client) {
	client = new(Client)
	client.options = defaultOptions()
	for _, opt := range opts {
		opt.apply(client.options)
	}
	client.aliyun = newAliyun()

	return
}

func (c *Client) Add(ctx context.Context, domain string, subdomain string, value string, opts ...option) (err error) {
	_options := c.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		err = c.aliyun.add(ctx, domain, subdomain, value, _options)
	}

	return
}

func (c *Client) Resolve(
	ctx context.Context,
	domain string, subdomain string, value string,
	opts ...option,
) (original string, do bool, err error) {
	_options := c.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	if record, getErr := c.Get(ctx, domain, subdomain, opts...); nil != getErr {
		err = getErr
	} else if nil != record && string(_options.typ) == record.Type {
		original = record.Value
		if value != record.Value {
			do = true
			err = c.Update(ctx, record, value, opts...)
		}
	} else {
		do = true
		if nil != record { // 先删除记录，不然会报重复错误
			_ = c.Delete(ctx, record, opts...)
		}
		err = c.Add(ctx, domain, subdomain, value, opts...)
	}

	return
}

func (c *Client) Get(ctx context.Context, domain string, subdomain string, opts ...option) (record *Record, err error) {
	_options := c.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		record, err = c.aliyun.get(ctx, domain, subdomain, _options)
	}

	return
}

func (c *Client) Update(ctx context.Context, record *Record, value string, opts ...option) (err error) {
	_options := c.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		err = c.aliyun.update(ctx, record, value, _options)
	}

	return
}

func (c *Client) Delete(ctx context.Context, record *Record, opts ...option) (err error) {
	_options := c.options.clone()
	for _, opt := range opts {
		opt.apply(_options)
	}

	switch _options.mode {
	case modeAliyun:
		err = c.aliyun.delete(ctx, record, _options)
	}

	return
}
