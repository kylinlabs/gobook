//第8章/code_8_6.go
package main

var pa [10000000]*int64

func main() {
	if err := trace.Start(os.Stdout); err != nil {
		log.Fatalln(err)
	}
	defer trace.Stop()
	for i := 0; i < 10000000; i++ {
		pa[i] = new(int64)
		pa[i] = new(int64)
	}
}
