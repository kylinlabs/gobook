//第7章/fence.go
package main

func main() {
	s := [2]chan struct{}{
		make(chan struct{}, 1),
		make(chan struct{}, 1),
	}
	f := make(chan struct{}, 2)
	var x, y, a, b int64
	go func() {
		for i := 0; i < 1000000; i++ {
			<-s[0]
			x = 1
			mfence()
			b = y
			f <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < 1000000; i++ {
			<-s[1]
			y = 1
			mfence()
			a = x
			f <- struct{}{}
		}
	}()
	for i := 0; i < 1000000; i++ {
		x = 0
		y = 0
		s[i%2] <- struct{}{}
		s[(i+1)%2] <- struct{}{}
		<-f
		<-f
		if a == 0 && b == 0 {
			println(i)
		}
	}
}

func mfence()
