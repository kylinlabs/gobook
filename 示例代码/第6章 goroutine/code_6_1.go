//第6章/code_6_1.go
package main

func hello(name string) {
	println("Hello ", name)
}
func main() {
	name := "Goroutine"
	go hello(name)
}
