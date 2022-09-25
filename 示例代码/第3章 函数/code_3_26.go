//第3章/code_3_26.go
func sc(n int) int {
	f := func() int {
		return n
	}
	n++
	return f()
}
