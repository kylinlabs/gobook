//第5章/code_5_17.go
func main() {
	n := Number(9)
	var a interface{} = n
	println(a)
	v := reflect.ValueOf("")
	v.MethodByName("")
}
