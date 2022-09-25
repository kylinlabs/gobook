//第3章/code_3_25.go
func sc(n int) int {
	n++
	f := func() int {
		return n
	}
	return f()
}
