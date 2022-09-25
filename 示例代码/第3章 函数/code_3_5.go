//第3章/code_3_5.go
package main

//go:noinline
func f1(a int8) (b int8) {
	println(&b, &a)
	return
}

func main() {
	f1(0)
}
