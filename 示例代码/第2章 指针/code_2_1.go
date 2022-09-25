//第2章/code_2_1.go
package main

func main() {
	n := 10
	println(read(&n))
}

//go:noinline
func read(p *int) (v int) {
	v = *p
	return
}
