//第3章/code_3_10.go
package main

func main() {
	println(*newInt())
}

//go:noinline
func newInt() *int {
	var a int
	return &a
}
