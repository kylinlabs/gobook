//第3章/code_3_12.go
var pt *int

//go:noinline
func setNew() {
	var a int
	pt = &a
}
