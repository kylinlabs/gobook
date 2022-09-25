//第3章/code_3_30.go
func fn(n int) (r int) {
	if n > 0 {
		defer func(i int) {
			r <<= i
		}(n)
	}
	n++
	return n
}
