//第2章/code_2_3.go
package main

var n int

func main() {
	println(addr())
}

//go:noinline
func addr() (p *int) {
	return &n
}
