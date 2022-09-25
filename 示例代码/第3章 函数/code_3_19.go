//第3章/code_3_19.go
func mc(n int) func() int {
	return func() int {
		return n
	}
}