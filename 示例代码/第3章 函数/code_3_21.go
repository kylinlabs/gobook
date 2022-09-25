//第3章/code_3_21.go
package main

import (
	"reflect"
	"unsafe"

	"github.com/fengyoulin/hookingo"
)

var hno hookingo.Hook

//go:linkname newobject runtime.newobject
func newobject(typ unsafe.Pointer) unsafe.Pointer

func fno(typ unsafe.Pointer) unsafe.Pointer {
	t := reflect.TypeOf(0)
	(*(*[2]unsafe.Pointer)(unsafe.Pointer(&t)))[1] = typ // 相当于反射了闭包对象类型
	println(t.String())
	if fn, ok := hno.Origin().(func(typ unsafe.Pointer) unsafe.Pointer); ok {
		return fn(typ) // 调用原runtime.newobject
	}
	return nil
}

// 创建一个闭包，make closure
func mc(start int) func() int {
	return func() int {
		start++
		return start
	}
}

func main() {
	var err error
	hno, err = hookingo.Apply(newobject, fno) // 应用钩子，替换函数
	if err != nil {
		panic(err)
	}
	f := mc(10)
	println(f())
}
