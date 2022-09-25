//第3章/code_3_13.go
var pp **int

//go:noinline
func dep() {
	var a int
	var p *int
	p = &a
	pp = &p
}
