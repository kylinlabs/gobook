//第3章/code_3_23.go
func sc(n int) int {
	f := func() int {
		return n
	}
	return f()
}
