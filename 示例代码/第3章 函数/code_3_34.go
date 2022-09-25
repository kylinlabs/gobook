//第3章/code_3_34.go
func fn() {
	defer func() {
		panic("2")
	}()
	panic("1")
}
