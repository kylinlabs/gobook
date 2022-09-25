//第3章/code_3_4.go
package main

type args struct {
	a int8
	b int64
	c int32
	d int16
}

//go:noinline
func f1(a args) (r args) {
	println(&r.d, &r.c, &r.b, &r.a, &a.d, &a.c, &a.b, &a.a)
	return
}

//go:noinline
func f2(aa int8, ab int64, ac int32, ad int16) (ra int8, rb int64, rc int32, rd int16) {
	println(&rd, &rc, &rb, &ra, &ad, &ac, &ab, &aa)
	return
}

func main() {
	f1(args{})
	f2(0, 0, 0, 0)
}
