//第8章/code_8_5.go
package main

var pa [10000000]**int64

func main() {
	for i := 0; i < 10000000; i++ {
		pa[i] = new(*int64)
		pa[i] = new(*int64)
	}
	time.Sleep(time.Second)
}
