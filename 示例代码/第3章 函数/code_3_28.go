//第3章/code_3_28.go
package main

func main() {
	println(df(10))
}

func df(n int) int {
	defer func(i *int) {
		*i *= 2
	}(&n)
	return n
}
