//第3章/code_3_27.go
func fn() func() {
	return func() {
		println("defer")
	}
}
