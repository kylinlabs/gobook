//第3章/code_3_16.go
//go:linkname retArg funny/inner.RetArg
func retArg(p *int) *int

//go:noinline
func arg() int {
	var a int
	var b int
	return *inner.RetArg(&a) + *retArg(&b)
}
