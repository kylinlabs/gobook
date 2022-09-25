//第5章/code_5_13.go
type Value struct {
	typ  unsafe.Pointer
	ptr  unsafe.Pointer
	flag uintptr
}

func toType(p unsafe.Pointer) (t reflect.Type) {
	t = reflect.TypeOf(0)
	(*[2]unsafe.Pointer)(unsafe.Pointer(&t))[1] = p
	return
}

func main() {
	n := 6789
	s := []interface{}{
		n,
		&n,
	}
	for i, v := range s {
		r := reflect.ValueOf(v)
		p := (*Value)(unsafe.Pointer(&r))
		println(i, p.typ, p.ptr, p.flag, toType(p.typ).String())
	}
}
