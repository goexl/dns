package uda

// Record 记录
type Record struct {
	// 编号
	Id string `json:"id"`
	// 域名
	Name string `json:"name"`
	// 子域名
	Subdomain string `json:"subdomain"`
	// 值
	Value string `json:"value"`
}
