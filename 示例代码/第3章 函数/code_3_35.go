//第3章/code_3_35.go
func fn() {
	defer func() {
		defer func() {
			recover()
		}
		panic("2")
	}()
	panic("1")
}
