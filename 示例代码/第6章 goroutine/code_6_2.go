//第6章/code_6_2.go
package main

import "fmt"

func main() {
	go func(n int) {
		for {
			n++
			fmt.Println(n)
		}
	}(0)
	for {
	}
}
