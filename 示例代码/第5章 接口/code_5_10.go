//第5章/code_5_10.go
package main

import (
	"fmt"
	"strconv"
	"unsafe"
)

type Integer int

func (i Integer) Value() float64 {
	return float64(i)
}
func (i Integer) String() string {
	return strconv.Itoa(int(i))
}

type Number interface {
	Value() float64
	String() string
}

func main() {
	i := Integer(0)
	fmt.Println(Methods(i))
	fmt.Println(Methods(&i))
	var n Number = i
	p := (*[5]unsafe.Pointer)((*face)(unsafe.Pointer(&n)).t)
	fmt.Println((*p)[3], (*p)[4])
}

func Methods(a interface{}) (r []Method) {
	e := (*face)(unsafe.Pointer(&a))
	u := uncommon(e.t)
	if u == nil {
		return nil
	}
	s := methods(u)
	r = make([]Method, len(s))
	for i := range s {
		r[i].Name = name(nameOff(e.t, s[i].name))
		r[i].Type = String(typeOff(e.t, s[i].mtyp))
		r[i].IFn = textOff(e.t, s[i].ifn)
		r[i].TFn = textOff(e.t, s[i].tfn)
	}
	return
}

type Method struct {
	Name string
	Type string
	IFn  unsafe.Pointer
	TFn  unsafe.Pointer
}

type face struct {
	t unsafe.Pointer
	d unsafe.Pointer
}

//go:linkname uncommon reflect.(*rtype).uncommon
func uncommon(t unsafe.Pointer) unsafe.Pointer

//go:linkname methods reflect.(*uncommonType).methods
func methods(u unsafe.Pointer) []method

type method struct {
	name int32
	mtyp int32
	ifn  int32
	tfn  int32
}

//go:linkname nameOff reflect.(*rtype).nameOff
func nameOff(t unsafe.Pointer, off int32) unsafe.Pointer

//go:linkname typeOff reflect.(*rtype).typeOff
func typeOff(t unsafe.Pointer, off int32) unsafe.Pointer

//go:linkname textOff reflect.(*rtype).textOff
func textOff(t unsafe.Pointer, off int32) unsafe.Pointer

//go:linkname String reflect.(*rtype).String
func String(t unsafe.Pointer) string

//go:linkname name reflect.name.name
func name(n unsafe.Pointer) string
