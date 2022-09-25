//第5章/code_5_11.go
package inherit

type A int

func (a A) Value() int {
	return int(a)
}

func (a *A) Set(n int) {
	*a = A(n)
}

type B struct {
	A
	b int
}

type C struct {
	*A
	c int
}
