//第5章/code_5_14.go
func main() {
	n := 6789
	v := reflect.ValueOf(&n)
	p := (*Value)(unsafe.Pointer(&v))
	println(p.typ, p.ptr, p.flag, toType(p.typ).String())
	e := v.Elem()
	p = (*Value)(unsafe.Pointer(&e))
	println(p.typ, p.ptr, p.flag, toType(p.typ).String())
	f := e.Addr()
	p = (*Value)(unsafe.Pointer(&f))
	println(p.typ, p.ptr, p.flag, toType(p.typ).String())
}
