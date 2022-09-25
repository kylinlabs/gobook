//第3章/code_3_22.go
func mc(n int) func() int {
	return func() int {
		return n
	}
}
