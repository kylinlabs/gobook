//第3章/code_3_6.go
//go:noinline
func fn() {
	var a int8
	var b int64
	var c int32
	var d int16
	var e int8
	println(&a, &b, &c, &d, &e)
}
