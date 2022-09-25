//第2章/code_2_5.go
//go:noinline
func convert(p *int) {
	q := (*int32)(unsafe.Pointer(p))
	*q = 0
}

