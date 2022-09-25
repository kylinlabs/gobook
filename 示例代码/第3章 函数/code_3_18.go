//第3章/code_3_18.go
package main

func main() {
	println(helper(nil, 0, 0))
}

//go:noinline
func helper(fn func(int, int) int, a, b int) int {
	return fn(a, b)
}
