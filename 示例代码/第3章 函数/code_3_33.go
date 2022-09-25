//第3章/code_3_33.go
func fn() {
	defer func() {
		r()
	}()
}

func r() {
	recover()
}
