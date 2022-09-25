//第3章/code_3_31.go
func fn() {
	defer func(a int) {
		recover()
		println(a)
	}(0)
}
