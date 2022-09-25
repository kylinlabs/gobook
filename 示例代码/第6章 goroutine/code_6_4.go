//第6章/code_6_4.go
package main
import "time"
func hello(name string){
	println("Hello ", name)
}
func main(){
	name := "Goroutine"
	go hello(name)
	time.Sleep(time.Second)
}
