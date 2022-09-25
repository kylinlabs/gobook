//第6章/code_6_3.go
func fibonacci(n int) int {
	if n < 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
