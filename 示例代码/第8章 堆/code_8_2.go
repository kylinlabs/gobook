//第8章/code_8_2.go
package main

var p *int64

func main() {
	for i := 0; i < 1000000; i++ {
		p = new(int64)
	}
	time.Sleep(time.Second)
}
