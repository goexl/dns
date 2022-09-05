package dns

var (
	_        = CNAME
	_        = A
	_        = AAAA
	_ option = (*optionType)(nil)
)

type optionType struct {
	typ Type
}

// CNAME CNAME解析
func CNAME() *optionType {
	return &optionType{
		typ: TypeCname,
	}
}

// A 地址解析
func A() *optionType {
	return &optionType{
		typ: TypeA,
	}
}

// AAAA 地址解析
func AAAA() *optionType {
	return &optionType{
		typ: TypeAAAA,
	}
}

func (t *optionType) apply(options *options) {
	options.typ = t.typ
}
