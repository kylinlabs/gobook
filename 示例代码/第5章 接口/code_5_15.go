//第5章/code_5_15.go
type Number float64

func (n Number) IntValue() int {
	return int(n)
}

func main() {
	n := Number(9)
	v := reflect.ValueOf(n)
	_ = v
}
