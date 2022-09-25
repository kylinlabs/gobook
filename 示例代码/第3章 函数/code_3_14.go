//第3章/code_3_14.go
package inner

//go:noinline
func RetArg(p *int) *int {
	return p
}
