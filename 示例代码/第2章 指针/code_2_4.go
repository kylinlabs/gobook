//第2章/code_2_4.go
//go:noinline
func newInt() (p *int) {
	var n int
	return &n
}
