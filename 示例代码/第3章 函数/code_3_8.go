//第3章/code_3_8.go
package main

func main() {
	in12(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12)
}

//go:noinline
func in12(a, b, c, d, e, f, g, h, i, j, k, l int8) int8 {
	return a + b + c + d + e + f + g + h + i + j + k + l
}
