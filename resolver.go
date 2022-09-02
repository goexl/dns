package uda

import (
	"context"
)

// Resolver 域名解析
type Resolver interface {
	// Add 添加解析记录
	Add(ctx context.Context, domain string, rr string, value string, opts ...option) (err error)

	// Resolve 域名解析
	Resolve(ctx context.Context, domain string, rr string, value string, opts ...option) (result *Result, err error)

	// Get 取域名记录
	Get(ctx context.Context, domain string, value string, opts ...option) (record *Record, err error)

	// Update 更新域名记录
	Update(ctx context.Context, record *Record, rr string, value string, opts ...option) (err error)
}
