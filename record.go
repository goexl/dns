package dns

import (
	"fmt"
)

// Record 记录
type Record struct {
	// 编号
	Id string `json:"id"`
	// 域名
	Name string `json:"name"`
	// 子域名
	Subdomain string `json:"subdomain"`
	// 类型
	Type string `json:"type"`
	// 值
	Value string `json:"value"`
}

func (r *Record) Final() string {
	return fmt.Sprintf(`%s.%s`, r.Name, r.Subdomain)
}
