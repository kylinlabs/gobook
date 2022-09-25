//第3章/code_3_32.go
func fn() {
	defer func() {
		recover()
	}()
}
