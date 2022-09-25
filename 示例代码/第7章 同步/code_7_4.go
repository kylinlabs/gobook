//第7章/code_7_4.go
package main

func main() {
	var wg sync.WaitGroup
	var x, y int64
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000000; i++ {
			if x == 0 {
				if y != 0 {
					println("1:", i)
				}
				x = 1
				y = 1
			}
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 1000000000; i++ {
			if y == 1 {
				if x != 1 {
					println("2:", i)
				}
				y = 0
				x = 0
			}
		}
	}()
	wg.Wait()
}
