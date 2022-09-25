//第4章/code_4_4.go
package gom

type A struct {
	a int
}

type B struct {
	b int
}

func (a A) Value() int {
	return a.a
}

func (a *A) Set(v int) {
	a.a = v
}

func (b B) Value() int {
	return b.b
}

func (b *B) Set(v int) {
	b.b = v
}
