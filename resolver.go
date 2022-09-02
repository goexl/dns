package uda

// Resolver 域名解析
type Resolver interface {
	// Resolve 域名解析
	Resolve(domain string, rr string, value string, opts ...option) (result Result, err error)
}
