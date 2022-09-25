//第5章/code_5_6.go
type A struct{}

type B struct{}

func (A) String() string {
	return "This is A"
}

func (B) String() string {
	return "This is B"
}

func toString(o fmt.Stringer) string {
	return o.String()
}

func main() {
	println(toString(A{})) // This is A
	println(toString(B{})) // This is B
}
