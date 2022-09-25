//第3章/code_3_15.go
package main

//go:noinline
func arg() int {
	var a int
	return *inner.RetArg(&a)
}
