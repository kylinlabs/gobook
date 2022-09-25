//第3章/code_3_29.go
func fn(n int) (r int) {
	if n&1 != 0 {
		defer func() {
			r <<= 1
		}()
	}
	for i := 0; i < n; i++ {
		defer func() {
			r <<= 1
		}()
	}
	return n
}
