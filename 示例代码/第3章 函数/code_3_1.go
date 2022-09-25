//第3章/code_3_1.go
package main

func main() {
	var v1, v2 int
	v3, v4 := f1(v1, v2)
	println(&v1, &v2, &v3, &v4)
	f2(v3)
}

//go:noinline
func f1(a1, a2 int) (r1, r2 int) {
	var l1, l2 int
	println(&r2, &r1, &a2, &a1, &l1, &l2)
	return
}

//go:noinline
func f2(a1 int) {
	println(&a1)
}
