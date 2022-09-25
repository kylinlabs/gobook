//第5章/code_5_16.go
func main() {
	n := Number(9)
	v := reflect.ValueOf(n)
	v.MethodByName("")
	_ = v
}
