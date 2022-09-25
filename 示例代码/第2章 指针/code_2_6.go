//第2章/code_2_6.go
//go:noinline
func size() (o uintptr) {
	o = unsafe.Sizeof(o)
	return
}
