package math

// nome com letra maiuscula para ser exportada (struct, func, etc)
type Math struct {
	A int
	B int
}

func (m Math) Add() int {
	return m.A + m.B
}

func (m Math) Sub() int {
	return m.A - m.B
}
