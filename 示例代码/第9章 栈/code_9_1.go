//第9章/code_9_1.go
//go:noinline
func test(i int) byte {
	var b [104]byte
	for x := range b {
		b[x] = byte(x)
	}
	return b[i%len(b)]
}
