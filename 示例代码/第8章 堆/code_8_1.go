//第8章/code_8_1.go
package main

var p *int

func main() {
	toStack(&p)
	toGlobal(&p)
	toUnknown(&p, &p)
}

//go:noinline
func toStack(i **int) (o *int) {
	o = *i
	return
}

//go:noinline
func toGlobal(i **int) {
	p = *i
}

//go:noinline
func toUnknown(a **int, i **int) {
	*a = *i
}
