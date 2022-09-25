//第3章/code_3_20.go
func mc2(a, b int) func() (int, int) {
	return func() (int, int) {
		return a, b
	}
}
