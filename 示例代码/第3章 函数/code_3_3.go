//第3章/code_3_3.go
package main

func main() {
	fa(0)
}

//go:noinline
func fa(n int) (r int) {
	r = fb(n)
	return
}

//go:noinline
func fb(n int) int {
	return n
}
