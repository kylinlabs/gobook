//第3章/code_3_11.go
//go:noinline
func New() int {
	p := new(int)
	return *p
}
