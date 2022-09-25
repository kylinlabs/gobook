//第5章/code_5_5.go
func fn(n int) bool {
	return notNil(n)
}

func notNil(a interface{}) bool {
	return a != nil
}
