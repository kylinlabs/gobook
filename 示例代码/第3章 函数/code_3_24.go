//第3章/code_3_24.go
func sc(n int) int {
	f := func() int {
		n++
		return n
	}
	return f()
}
