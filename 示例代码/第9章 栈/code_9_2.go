//第9章/code_9_2.go
func mc(l int) func(i int) byte {
	return func(i int) byte {
		b := make([]byte, l)
		for x := range b {
			b[x] = byte(x)
		}
		return b[i%len(b)]
	}
}
